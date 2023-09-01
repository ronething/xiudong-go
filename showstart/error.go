package showstart

import (
	"fmt"
)

var _ error = &ResponseWrapper{}

type ResponseWrapper struct {
	Success bool        `json:"success"`
	State   interface{} `json:"state"`
	Msg     string      `json:"msg"`
}

func (s *ResponseWrapper) Error() string {
	return fmt.Sprintf("msg: %v, success: %v, state: %v", s.Msg, s.Success, s.State)
}
