package teq

import "regexp"

// IsEmail check a string is an email or not
func IsEmail(str string) (bool, error) {
	match, err := regexp.MatchString(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, str)
	if err != nil {
		return false, err
	}

	return match, nil
}
