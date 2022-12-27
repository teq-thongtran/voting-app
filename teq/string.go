package teq

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// ConvertToNonAccentVietnamese convert a string into a string has no accent Vietnamese
func ConvertToNonAccentVietnamese(text string) string {
	type replace struct {
		regex string
		value string
	}

	replaces := []replace{
		{
			regex: "à|á|ạ|ả|ã|â|ầ|ấ|ậ|ẩ|ẫ|ă|ằ|ắ|ặ|ẳ|ẵ",
			value: "a",
		},
		{
			regex: "è|é|ẹ|ẻ|ẽ|ê|ề|ế|ệ|ể|ễ",
			value: "e",
		},
		{
			regex: "ì|í|ị|ỉ|ĩ",
			value: "i",
		},
		{
			regex: "ò|ó|ọ|ỏ|õ|ô|ồ|ố|ộ|ổ|ỗ|ơ|ờ|ớ|ợ|ở|ỡ",
			value: "o",
		},
		{
			regex: "ù|ú|ụ|ủ|ũ|ư|ừ|ứ|ự|ử|ữ",
			value: "u",
		},
		{
			regex: "ỳ|ý|ỵ|ỷ|ỹ",
			value: "y",
		},
		{
			regex: "đ",
			value: "d",
		},
	}

	for i := range replaces {
		re := regexp.MustCompile(replaces[i].regex)
		text = re.ReplaceAllString(text, replaces[i].value)

		re = regexp.MustCompile(strings.ToUpper(replaces[i].regex))
		text = re.ReplaceAllString(text, strings.ToUpper(replaces[i].value))
	}

	return text
}

// RandomString random a string based on the configuration parameters that you specified
func RandomString(length int, hasSpecialCharacter bool) string {
	rand.Seed(time.Now().UnixNano())

	var (
		buf      = make([]byte, length)
		digits   = "0123456789"
		specials = "~=+%^*/()[]{}/!@#$?|"
		all      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + digits
		temp     = 2
	)

	buf[0] = digits[rand.Intn(len(digits))]

	if hasSpecialCharacter {
		all += specials
		buf[1] = specials[rand.Intn(len(specials))]
	} else {
		temp--
	}

	for i := temp; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}
