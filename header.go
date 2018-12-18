package main

import "fmt"

func BuildHeader(guid string, clientId string) string {
	result := "<Header>"
	result += fmt.Sprintf("<ClientId>%s</ClientId>", clientId)
	result += fmt.Sprintf("<Guid>%s</Guid>", guid)
	result += "</Header>"
	return result
}
