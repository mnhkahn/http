package http

var DEFAULT_CONTROLLER *Controller = new(Controller)

type ControllerIfac interface {
	Init(ctx *Context)
}

type Controller struct {
	Ctx *Context
}

func (this *Controller) Init(ctx *Context) {
	this.Ctx = ctx
}
