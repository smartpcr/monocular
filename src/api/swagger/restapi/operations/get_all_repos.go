package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllReposHandlerFunc turns a function with the right signature into a get all repos handler
type GetAllReposHandlerFunc func(GetAllReposParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllReposHandlerFunc) Handle(params GetAllReposParams) middleware.Responder {
	return fn(params)
}

// GetAllReposHandler interface for that can handle valid get all repos params
type GetAllReposHandler interface {
	Handle(GetAllReposParams) middleware.Responder
}

// NewGetAllRepos creates a new http.Handler for the get all repos operation
func NewGetAllRepos(ctx *middleware.Context, handler GetAllReposHandler) *GetAllRepos {
	return &GetAllRepos{Context: ctx, Handler: handler}
}

/*GetAllRepos swagger:route GET /v1/repos getAllRepos

get all repositories enabled in the backend

*/
type GetAllRepos struct {
	Context *middleware.Context
	Handler GetAllReposHandler
}

func (o *GetAllRepos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetAllReposParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
