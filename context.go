package http

import (
//	"time"
)

type Context struct {
	Req  *Request
	Resp *Response
	//	elapse time.Duration
}

func NewContext() *Context {
	ctx := new(Context)
	return ctx
}
