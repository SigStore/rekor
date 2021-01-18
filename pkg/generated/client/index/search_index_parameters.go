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

package index

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

	"github.com/projectrekor/rekor/pkg/generated/models"
)

// NewSearchIndexParams creates a new SearchIndexParams object
// with the default values initialized.
func NewSearchIndexParams() *SearchIndexParams {
	var ()
	return &SearchIndexParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchIndexParamsWithTimeout creates a new SearchIndexParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchIndexParamsWithTimeout(timeout time.Duration) *SearchIndexParams {
	var ()
	return &SearchIndexParams{

		timeout: timeout,
	}
}

// NewSearchIndexParamsWithContext creates a new SearchIndexParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchIndexParamsWithContext(ctx context.Context) *SearchIndexParams {
	var ()
	return &SearchIndexParams{

		Context: ctx,
	}
}

// NewSearchIndexParamsWithHTTPClient creates a new SearchIndexParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchIndexParamsWithHTTPClient(client *http.Client) *SearchIndexParams {
	var ()
	return &SearchIndexParams{
		HTTPClient: client,
	}
}

/*SearchIndexParams contains all the parameters to send to the API endpoint
for the search index operation typically these are written to a http.Request
*/
type SearchIndexParams struct {

	/*Query*/
	Query *models.SearchIndex

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search index params
func (o *SearchIndexParams) WithTimeout(timeout time.Duration) *SearchIndexParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search index params
func (o *SearchIndexParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search index params
func (o *SearchIndexParams) WithContext(ctx context.Context) *SearchIndexParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search index params
func (o *SearchIndexParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search index params
func (o *SearchIndexParams) WithHTTPClient(client *http.Client) *SearchIndexParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search index params
func (o *SearchIndexParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithQuery adds the query to the search index params
func (o *SearchIndexParams) WithQuery(query *models.SearchIndex) *SearchIndexParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search index params
func (o *SearchIndexParams) SetQuery(query *models.SearchIndex) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *SearchIndexParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Query != nil {
		if err := r.SetBodyParam(o.Query); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
