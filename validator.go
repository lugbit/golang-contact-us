package main

import (
	"regexp"
	"strings"
)

type formValidate struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Subject   string
	Body      string
	Errors    map[string]string
	SentFlag  bool
}

func (fv *formValidate) validate() bool {
	// Map will not be empty if errors are present.
	fv.Errors = make(map[string]string)
	// If the first name is empty, add to errors map.
	if strings.TrimSpace(fv.FirstName) == "" {
		fv.Errors["FirstName"] = "First name cannot be empty."
	}
	if strings.TrimSpace(fv.LastName) == "" {
		fv.Errors["LastName"] = "Last name cannot be empty."
	}
	// If email is not empty, check if it's the correct format.
	if strings.TrimSpace(fv.Email) != "" {
		re := regexp.MustCompile(".+@.+\\..+")
		matched := re.Match([]byte(fv.Email))

		if matched == false {
			fv.Errors["EmailFormat"] = "Email is not valid."
		}
	} else {
		fv.Errors["Email"] = "Email cannot be empty."
	}

	if strings.TrimSpace(fv.Subject) == "" {
		fv.Errors["Subject"] = "Subject cannot be empty."
	}
	if strings.TrimSpace(fv.Body) == "" {
		fv.Errors["Body"] = "Body cannot be empty."
	}
	// Return empty map (assuming it's empty/no errors)
	return len(fv.Errors) == 0
}
