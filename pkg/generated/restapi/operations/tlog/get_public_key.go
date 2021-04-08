// Code generated by go-swagger; DO NOT EDIT.

// /*
// Copyright The Rekor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPublicKeyHandlerFunc turns a function with the right signature into a get public key handler
type GetPublicKeyHandlerFunc func(GetPublicKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPublicKeyHandlerFunc) Handle(params GetPublicKeyParams) middleware.Responder {
	return fn(params)
}

// GetPublicKeyHandler interface for that can handle valid get public key params
type GetPublicKeyHandler interface {
	Handle(GetPublicKeyParams) middleware.Responder
}

// NewGetPublicKey creates a new http.Handler for the get public key operation
func NewGetPublicKey(ctx *middleware.Context, handler GetPublicKeyHandler) *GetPublicKey {
	return &GetPublicKey{Context: ctx, Handler: handler}
}

/* GetPublicKey swagger:route GET /api/v1/log/publicKey tlog getPublicKey

Retrieve the public key that can be used to validate the signed tree head

Returns the public key that can be used to validate the signed tree head

*/
type GetPublicKey struct {
	Context *middleware.Context
	Handler GetPublicKeyHandler
}

func (o *GetPublicKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetPublicKeyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
