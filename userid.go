// A simple UserID structure with field validation
// Validation is performed by using regular expressions to validate characters in allow lists
// i.e. for a "name" field, we validate by only allowing certain characters.
// It is a common mistake to use block list validation in order to try to detect possibly
// dangerous characters and patterns like the apostrophe ' character, the string 1=1, or the <script> tag, but
// this is a massively flawed approach as it is trivial for an attacker to bypass such filters.

// OWASP Input Validation Cheat Sheet
// For reference: https://cheatsheetseries.owasp.org/cheatsheets/Input_Validation_Cheat_Sheet.html#allow-list-vs-block-list

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	idRegex           = "^[a-zA-Z0-9_.-]*$"
	nameRegex         = "^([A-Za-z]+[-' ]?)*[A-Za-z]$"
	mobileNumberRegex = "^04\\d{7,9}$"
)

var (
	idMatcher           = regexp.MustCompile(idRegex)
	nameMatcher         = regexp.MustCompile(nameRegex)
	mobileNumberMatcher = regexp.MustCompile(mobileNumberRegex)
)

type UserID struct {
	Name         string
	MobileNumber string
	MiddleName   string
	LastName     string 
}

func (a *UserID) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Name, validateAndLogError("Name", validation.Required, validation.Length(1, 25), validation.Match(nameMatcher))),
		validation.Field(&a.MobileNumber, validateAndLogError("MobileNumber", validation.Required, validation.Length(6, 10), validation.Match(mobileNumberMatcher))),
		validation.Field(&a.MiddleName, validateAndLogError("MiddleName", validation.Required ))),
		validation.Field(&a.LastNamer, validateAndLogError("LastName",  validation.Length(6, 10))
	)
}

func ValidateID(value interface{}) error {
	id, ok := value.(*string)
	if !ok {
		return errors.New("invalid id")
	}
	return validation.Validate(id, validateAndLogError("ID", validation.Length(36, 36), validation.Match(idMatcher)))
}

func validateAndLogError(fieldName string, rules ...validation.Rule) validation.Rule {
	return validation.WithContext(func(ctx context.Context, value interface{}) error {
		var err = validation.Validate(value, rules...)
		if err != nil {
			log.Println(fmt.Sprintf("invalid: '%s'. error: '%s'.", fieldName, err.Error()))
		}
		return err
	})
}
