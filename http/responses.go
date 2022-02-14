package http

import "stvcv2/commandFormatter"

type Response struct {
	Message string
	Data    []commandFormatter.CommandResult
	Code    int
}
