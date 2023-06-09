package routes 

type Route struct {
	routeType  string
	routeParam string
}

func (r *Route) addRoute(routeType string , routeParam string){
  r.routeParam = routeParam
  r.routeType  = routeType
}
