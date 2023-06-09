package router 

type Route struct {
	routeType  string
	routeParam string
}

var Routes []Route

func New(rt string, rp string) *Route {
	return &Route{routeType: rt, routeParam: rp}
}

func (r *Route) addRoute(routeType string, routeParam string) {
	rt := routeType
	rp := routeParam
	route := New(rt, rp)
	Routes = append(Routes, *route)
}

func (r *Route) Routes() []Route {
	return Routes
}
