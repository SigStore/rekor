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

package pubkey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pubkey API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pubkey API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetPublicKey(params *GetPublicKeyParams, opts ...ClientOption) (*GetPublicKeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetPublicKey retrieves the public key that can be used to validate the signed tree head

  Returns the public key that can be used to validate the signed tree head
*/
func (a *Client) GetPublicKey(params *GetPublicKeyParams, opts ...ClientOption) (*GetPublicKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPublicKey",
		Method:             "GET",
		PathPattern:        "/api/v1/log/publicKey",
		ProducesMediaTypes: []string{"application/x-pem-file"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPublicKeyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPublicKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetPublicKeyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
