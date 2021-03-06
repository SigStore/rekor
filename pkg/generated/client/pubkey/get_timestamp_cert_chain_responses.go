// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
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
//

package pubkey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sigstore/rekor/pkg/generated/models"
)

// GetTimestampCertChainReader is a Reader for the GetTimestampCertChain structure.
type GetTimestampCertChainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTimestampCertChainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTimestampCertChainOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTimestampCertChainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTimestampCertChainOK creates a GetTimestampCertChainOK with default headers values
func NewGetTimestampCertChainOK() *GetTimestampCertChainOK {
	return &GetTimestampCertChainOK{}
}

/* GetTimestampCertChainOK describes a response with status code 200, with default header values.

The PEM encoded cert chain
*/
type GetTimestampCertChainOK struct {
	Payload string
}

func (o *GetTimestampCertChainOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/timestampCertChain][%d] getTimestampCertChainOK  %+v", 200, o.Payload)
}
func (o *GetTimestampCertChainOK) GetPayload() string {
	return o.Payload
}

func (o *GetTimestampCertChainOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTimestampCertChainDefault creates a GetTimestampCertChainDefault with default headers values
func NewGetTimestampCertChainDefault(code int) *GetTimestampCertChainDefault {
	return &GetTimestampCertChainDefault{
		_statusCode: code,
	}
}

/* GetTimestampCertChainDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetTimestampCertChainDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get timestamp cert chain default response
func (o *GetTimestampCertChainDefault) Code() int {
	return o._statusCode
}

func (o *GetTimestampCertChainDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/timestampCertChain][%d] getTimestampCertChain default  %+v", o._statusCode, o.Payload)
}
func (o *GetTimestampCertChainDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTimestampCertChainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
