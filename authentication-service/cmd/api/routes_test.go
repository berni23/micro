package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	
)

//test routes actually exist
func Test_routes_exist(t *testing.T){

	testApp:= Config {}

	testRoutes:= testApp.routes() //returns a http handler

	chiRoutes:= testRoutes.(chi.Router)
	
	//one entry for each route a particular microservice has
	routes:=[]string{"/authenticate"}


	for _,route:=range routes {
		routeExists(t,chiRoutes,route)
	}
	

}



func routeExists(t *testing.T,routes chi.Router, route string){
	found:=false 
	_=chi.Walk(routes,func(method string,foundRoute string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		if route==foundRoute {
			found = true 
			return nil
		}

		return nil

	})


	if  !found{

		t.Errorf("did not find %s in registered routes",route)

	}


}