package main

import "strings"

/*
 * strings.HasPrefix() API, templateStr--
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	templateStr := strs[0]
	for i := 1; len(templateStr) > 0 && i < len(strs); {
		if strings.HasPrefix(strs[i], templateStr) {
			i++
			continue
		}
		templateStr = templateStr[:len(templateStr)-1]

	}
	return templateStr
}
