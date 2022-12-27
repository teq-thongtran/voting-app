package teq

import "regexp"

// IsTIMEw check a string is right TIMEw or not
// TIMEw is format time HH:mm:ss.
func IsTIMEw(str string) (bool, error) {
	match, err := regexp.MatchString("^(?:[01]\\d|2[0123]):(?:[012345]\\d):(?:[012345]\\d)$", str)
	if err != nil {
		return false, err
	}

	return match, nil
}
