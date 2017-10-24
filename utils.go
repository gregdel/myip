package main

import "strings"

func splitIP(s string) string {
	return strings.Split(s, ":")[0]
}

func lastIndexIP(s string) string {
	return s[:strings.LastIndex(s, ":")]
}
