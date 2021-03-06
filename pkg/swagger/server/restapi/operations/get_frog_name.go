// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFrogNameHandlerFunc turns a function with the right signature into a get frog name handler
type GetFrogNameHandlerFunc func(GetFrogNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFrogNameHandlerFunc) Handle(params GetFrogNameParams) middleware.Responder {
	return fn(params)
}

// GetFrogNameHandler interface for that can handle valid get frog name params
type GetFrogNameHandler interface {
	Handle(GetFrogNameParams) middleware.Responder
}

// NewGetFrogName creates a new http.Handler for the get frog name operation
func NewGetFrogName(ctx *middleware.Context, handler GetFrogNameHandler) *GetFrogName {
	return &GetFrogName{Context: ctx, Handler: handler}
}

/* GetFrogName swagger:route GET /frog/{name} getFrogName

Return the frog Image.

*/
type GetFrogName struct {
	Context *middleware.Context
	Handler GetFrogNameHandler
}

func (o *GetFrogName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFrogNameParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
