package main

import "fmt"

func BuildHeader(guid string, clientId int) string {
	result := "<Header>"
	result += fmt.Sprintf("<ClientId>%d</ClientId>", clientId)
	result += fmt.Sprintf("<Guid>%s</Guid>", guid)
	result += "</Header>"
	return result
}
