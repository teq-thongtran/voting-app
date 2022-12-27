package teq

import "strings"

func SQLEscapeString(val string) string {
	replacer := strings.NewReplacer(
		"\\0", "\\\\0",
		"\n", "\\n",
		"\r", "\\r",
		"\x1a", "\\Z",
		`"`, `\"`,
		"'", `\'`,
		"\\", "\\\\",
	)

	return replacer.Replace(val)
}
