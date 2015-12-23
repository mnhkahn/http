package http

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var DEFAULT_CONTROLLER *Controller = new(Controller)

type ControllerIfac interface {
	Init(ctx *Context)
	Finish()
}

type Controller struct {
	Ctx       *Context
	TemplPath string
}

func (this *Controller) Init(ctx *Context) {
	this.Ctx = ctx
}

func (this *Controller) Option() {
	allowMethods := []string{}
	for _, method := range HTTP_METHOD {
		allowMethods = append(allowMethods, method)
	}
	this.Ctx.Resp.Headers.Add(HTTP_HEAD_ALLOW, strings.Join(allowMethods, ", "))
}

func (this *Controller) ServeJson(j interface{}) {
	this.Ctx.Resp.Headers.Add(HTTP_HEAD_CONTENTTYPE, "text/plain; charset=utf-8")
	v, _ := json.Marshal(j)
	this.Ctx.Resp.Body = string(v)
}

func (this *Controller) ServeView(params ...interface{}) {
	if len(params) <= 0 {

	} else if len(params) == 1 {
		this.Ctx.Resp.Headers.Add(HTTP_HEAD_CONTENTTYPE, "text/html; charset=utf-8")
		if templ, exists := ViewsTemplFiles[params[0].(string)]; exists {
			v, _ := ioutil.ReadFile(templ)
			this.Ctx.Resp.Body = string(v)
		} else {
			ErrLog.Println("Can't find the template file", params)
		}
	} else {

	}
}

func (this *Controller) ServeFile(params ...interface{}) {
	if len(params) <= 0 {

	} else if len(params) == 1 {
		if templ, exists := ViewsTemplFiles[params[0].(string)]; exists {
			this.Ctx.Resp.Body = string(templ)
		} else {
			ErrLog.Println("Can't find the template file", params)
		}
	} else {

	}
}

func (this *Controller) Favicon() {
	this.Ctx.Resp.StatusCode = StatusFound
	this.Ctx.Resp.Headers.Add(HTTP_HEAD_LOCATION, "http://7b1h1l.com1.z0.glb.clouddn.com/c32.ico")
}

func (this *Controller) Finish() {
}
