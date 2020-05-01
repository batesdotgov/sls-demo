package helpers

import (
	"errors"
	"net/url"
)

// ParseAndCheckBody check
func ParseAndCheckBody(requestBody string) error {
	values, err := url.ParseQuery(requestBody)

	if err != nil {
		return errors.New("could not parse request body")
	}

	name := values.Get("name")

	if name == "" {
		return errors.New("name value required")
	}

	return nil
}
