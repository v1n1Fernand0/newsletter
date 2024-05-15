package formatter

import "strings"

func FormatNewsletter(content []string) string {
	return strings.Join(content, "\n\n")
}
