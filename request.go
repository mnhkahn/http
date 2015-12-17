package http

import (
	"strings"
)

type Request struct {
	Method string

	Url string

	Proto string

	UserAgent string

	Host string

	Headers Header

	Body string
}

func NewRequst(b string) *Request {
	r := new(Request)
	r.Headers = make(map[string][]string, 0)

	startLine := strings.Split(b[:strings.Index(b, CRLF)], " ")
	if len(startLine) == 3 {
		r.Method, r.Url, r.Proto = startLine[0], startLine[1], startLine[2]
	}
	b = b[strings.Index(b, CRLF)+len(CRLF):]

	if strings.LastIndex(b, CRLF+CRLF) != -1 {
		r.Body = b[strings.LastIndex(b, CRLF+CRLF):]
		b = b[:strings.LastIndex(b, CRLF+CRLF)-2]
	}
	b = strings.TrimSpace(b)

	for _, line := range strings.Split(b, CRLF) {
		k, v := line[:strings.Index(line, ":")], line[strings.Index(line, ":")+1:]
		k, v = strings.TrimSpace(k), strings.TrimSpace(v)
		if k == HTTP_HEAD_USERAGENT {
			r.UserAgent = v
		} else if k == HTTP_HEAD_HOST {
			r.Host = v
		} else {
			r.Headers[k] = append(r.Headers[k], v)
		}
	}

	return r
}
