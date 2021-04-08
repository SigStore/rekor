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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetPublicKeyCertParams creates a new GetPublicKeyCertParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPublicKeyCertParams() *GetPublicKeyCertParams {
	return &GetPublicKeyCertParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicKeyCertParamsWithTimeout creates a new GetPublicKeyCertParams object
// with the ability to set a timeout on a request.
func NewGetPublicKeyCertParamsWithTimeout(timeout time.Duration) *GetPublicKeyCertParams {
	return &GetPublicKeyCertParams{
		timeout: timeout,
	}
}

// NewGetPublicKeyCertParamsWithContext creates a new GetPublicKeyCertParams object
// with the ability to set a context for a request.
func NewGetPublicKeyCertParamsWithContext(ctx context.Context) *GetPublicKeyCertParams {
	return &GetPublicKeyCertParams{
		Context: ctx,
	}
}

// NewGetPublicKeyCertParamsWithHTTPClient creates a new GetPublicKeyCertParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPublicKeyCertParamsWithHTTPClient(client *http.Client) *GetPublicKeyCertParams {
	return &GetPublicKeyCertParams{
		HTTPClient: client,
	}
}

/* GetPublicKeyCertParams contains all the parameters to send to the API endpoint
   for the get public key cert operation.

   Typically these are written to a http.Request.
*/
type GetPublicKeyCertParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get public key cert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPublicKeyCertParams) WithDefaults() *GetPublicKeyCertParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get public key cert params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPublicKeyCertParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get public key cert params
func (o *GetPublicKeyCertParams) WithTimeout(timeout time.Duration) *GetPublicKeyCertParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get public key cert params
func (o *GetPublicKeyCertParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get public key cert params
func (o *GetPublicKeyCertParams) WithContext(ctx context.Context) *GetPublicKeyCertParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get public key cert params
func (o *GetPublicKeyCertParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get public key cert params
func (o *GetPublicKeyCertParams) WithHTTPClient(client *http.Client) *GetPublicKeyCertParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get public key cert params
func (o *GetPublicKeyCertParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicKeyCertParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}