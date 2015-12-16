package http


// /favicon.ico http://7b1h1l.com1.z0.glb.clouddn.com/c32.ico
type Route struct {
	routes map[string]map[string]Handler
}

func NewRoute() *Route {
	r := new(Route)
	r.routes = make(map[string]map[string]Handler)
	return r
}
