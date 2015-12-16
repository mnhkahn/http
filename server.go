package http

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"reflect"
	"strings"
	"time"
)

var HTTP_METHOD = map[string]string{
	"GET":     "GET",
	"POST":    "POST",
	"HEAD":    "HEAD",
	"PUT":     "PUT",
	"TRACE":   "TRACE",
	"OPTIONS": "OPTIONS",
	"DELETE":  "DELETE",
}

type Address struct {
	Host string
	Port string
}

func NewAddress(addr_str string) *Address {
	addr := new(Address)
	addr_strs := strings.Split(addr_str, ":")
	if len(addr_strs) == 2 {
		addr.Host, addr.Port = addr_strs[0], addr_strs[1]
	}
	return addr
}

func (this *Address) String() string {
	return this.Host + ":" + this.Port
}

type Server struct {
	Addr   *Address
	Routes *Route
}

var DEFAULT_SERVER *Server

func Serve(addr string) {
	DEFAULT_SERVER.Addr = NewAddress(addr)
	ln, err := net.Listen("tcp", DEFAULT_SERVER.Addr.String())
	defer ln.Close()
	if err != nil {
		panic(err)
	}

	log.Printf("<<<Server Accepting on Port %s>>>\n\n", DEFAULT_SERVER.Addr.Port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go handleConnection(conn)
	}
}

func init() {
	DEFAULT_SERVER = new(Server)
	DEFAULT_SERVER.Routes = NewRoute()
}

func handleConnection(conn net.Conn) {
	serve_time := time.Now()

	defer conn.Close()

	ctx := NewContext()
	ctx.Resp = new(Response)

	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		log.Println("Error to read message because of ", err)
		ctx.Resp.StatusCode = StatusInternalServerError
		goto END
	}
	ctx.Req = NewRequst(string(buf[:reqLen-1]))
	ctx.Resp.Proto = ctx.Req.Proto
	if DEFAULT_SERVER.Routes.routes[ctx.Req.Method][ctx.Req.Url] != nil {
		DEFAULT_SERVER.Routes.routes[ctx.Req.Method][ctx.Req.Url].ServeHTTP(ctx)
		ctx.Resp.StatusCode = StatusOK
	} else {
		panic(2)
		ctx.Resp.StatusCode = StatusNotFound
	}

END:

	if ctx.Resp.StatusCode == StatusNotFound {
		ctx.Resp.Body = DEFAULT_ERROR_PAGE
	}
	buffers := bytes.Buffer{}
	buffers.WriteString(fmt.Sprintf("%s %d %s\r\n", ctx.Resp.Proto, ctx.Resp.StatusCode, StatusText(ctx.Resp.StatusCode)))
	buffers.WriteString("Server: Cyeam\r\n")
	buffers.WriteString("Date: " + serve_time.Format(time.RFC1123) + "\r\n")
	buffers.WriteString("Content-Type: text/html; charset=utf-8\r\n")
	buffers.WriteString("Content-length:" + fmt.Sprintf("%d", len(ctx.Resp.Body)) + "\r\n")
	buffers.WriteString("\r\n")
	buffers.WriteString(ctx.Resp.Body)
	_, err = conn.Write(buffers.Bytes())
	if err != nil {
		log.Println(err)
	}
	log.Println(ctx.Resp.StatusCode, ctx.Req.Method, ctx.Req.Url, ctx.Req.UserAgent, conn.RemoteAddr(), time.Now().Sub(serve_time))
}

func Router(path string, method string, ctrl ControllerIfac, methodName string) {
	r := make(map[string]Handler)
	handler := new(Handle)
	handler.ctrl = ctrl
	handler.methodName = methodName
	handler.fn = reflect.ValueOf(handler.ctrl).MethodByName(handler.methodName)
	r[path] = handler
	DEFAULT_SERVER.Routes.routes[method] = r
}

var DEFAULT_ERROR_PAGE = "<iframe scrolling='no' frameborder='0' src='http://yibo.iyiyun.com/js/yibo404/key/2354' width='640' height='464' style='display:block;'></iframe>"
