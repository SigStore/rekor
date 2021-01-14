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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/runtime/yamlpc"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/projectrekor/rekor/pkg/generated/restapi/operations/entries"
	"github.com/projectrekor/rekor/pkg/generated/restapi/operations/tlog"
)

// NewRekorServerAPI creates a new RekorServer instance
func NewRekorServerAPI(spec *loads.Document) *RekorServerAPI {
	return &RekorServerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),
		YamlConsumer: yamlpc.YAMLConsumer(),

		ApplicationXPemFileProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("applicationXPemFile producer has not yet been implemented")
		}),
		JSONProducer: runtime.JSONProducer(),
		YamlProducer: yamlpc.YAMLProducer(),

		EntriesCreateLogEntryHandler: entries.CreateLogEntryHandlerFunc(func(params entries.CreateLogEntryParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.CreateLogEntry has not yet been implemented")
		}),
		EntriesGetLogEntryByIndexHandler: entries.GetLogEntryByIndexHandlerFunc(func(params entries.GetLogEntryByIndexParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.GetLogEntryByIndex has not yet been implemented")
		}),
		EntriesGetLogEntryByUUIDHandler: entries.GetLogEntryByUUIDHandlerFunc(func(params entries.GetLogEntryByUUIDParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.GetLogEntryByUUID has not yet been implemented")
		}),
		EntriesGetLogEntryProofHandler: entries.GetLogEntryProofHandlerFunc(func(params entries.GetLogEntryProofParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.GetLogEntryProof has not yet been implemented")
		}),
		TlogGetLogInfoHandler: tlog.GetLogInfoHandlerFunc(func(params tlog.GetLogInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation tlog.GetLogInfo has not yet been implemented")
		}),
		TlogGetLogProofHandler: tlog.GetLogProofHandlerFunc(func(params tlog.GetLogProofParams) middleware.Responder {
			return middleware.NotImplemented("operation tlog.GetLogProof has not yet been implemented")
		}),
		TlogGetPublicKeyHandler: tlog.GetPublicKeyHandlerFunc(func(params tlog.GetPublicKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation tlog.GetPublicKey has not yet been implemented")
		}),
		EntriesSearchLogQueryHandler: entries.SearchLogQueryHandlerFunc(func(params entries.SearchLogQueryParams) middleware.Responder {
			return middleware.NotImplemented("operation entries.SearchLogQuery has not yet been implemented")
		}),
	}
}

/*RekorServerAPI Rekor is a cryptographically secure, immutable transparency log for signed software releases. */
type RekorServerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// YamlConsumer registers a consumer for the following mime types:
	//   - application/yaml
	YamlConsumer runtime.Consumer

	// ApplicationXPemFileProducer registers a producer for the following mime types:
	//   - application/x-pem-file
	ApplicationXPemFileProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// YamlProducer registers a producer for the following mime types:
	//   - application/yaml
	YamlProducer runtime.Producer

	// EntriesCreateLogEntryHandler sets the operation handler for the create log entry operation
	EntriesCreateLogEntryHandler entries.CreateLogEntryHandler
	// EntriesGetLogEntryByIndexHandler sets the operation handler for the get log entry by index operation
	EntriesGetLogEntryByIndexHandler entries.GetLogEntryByIndexHandler
	// EntriesGetLogEntryByUUIDHandler sets the operation handler for the get log entry by UUID operation
	EntriesGetLogEntryByUUIDHandler entries.GetLogEntryByUUIDHandler
	// EntriesGetLogEntryProofHandler sets the operation handler for the get log entry proof operation
	EntriesGetLogEntryProofHandler entries.GetLogEntryProofHandler
	// TlogGetLogInfoHandler sets the operation handler for the get log info operation
	TlogGetLogInfoHandler tlog.GetLogInfoHandler
	// TlogGetLogProofHandler sets the operation handler for the get log proof operation
	TlogGetLogProofHandler tlog.GetLogProofHandler
	// TlogGetPublicKeyHandler sets the operation handler for the get public key operation
	TlogGetPublicKeyHandler tlog.GetPublicKeyHandler
	// EntriesSearchLogQueryHandler sets the operation handler for the search log query operation
	EntriesSearchLogQueryHandler entries.SearchLogQueryHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *RekorServerAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *RekorServerAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *RekorServerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *RekorServerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *RekorServerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *RekorServerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *RekorServerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *RekorServerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *RekorServerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the RekorServerAPI
func (o *RekorServerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.YamlConsumer == nil {
		unregistered = append(unregistered, "YamlConsumer")
	}

	if o.ApplicationXPemFileProducer == nil {
		unregistered = append(unregistered, "ApplicationXPemFileProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.YamlProducer == nil {
		unregistered = append(unregistered, "YamlProducer")
	}

	if o.EntriesCreateLogEntryHandler == nil {
		unregistered = append(unregistered, "entries.CreateLogEntryHandler")
	}
	if o.EntriesGetLogEntryByIndexHandler == nil {
		unregistered = append(unregistered, "entries.GetLogEntryByIndexHandler")
	}
	if o.EntriesGetLogEntryByUUIDHandler == nil {
		unregistered = append(unregistered, "entries.GetLogEntryByUUIDHandler")
	}
	if o.EntriesGetLogEntryProofHandler == nil {
		unregistered = append(unregistered, "entries.GetLogEntryProofHandler")
	}
	if o.TlogGetLogInfoHandler == nil {
		unregistered = append(unregistered, "tlog.GetLogInfoHandler")
	}
	if o.TlogGetLogProofHandler == nil {
		unregistered = append(unregistered, "tlog.GetLogProofHandler")
	}
	if o.TlogGetPublicKeyHandler == nil {
		unregistered = append(unregistered, "tlog.GetPublicKeyHandler")
	}
	if o.EntriesSearchLogQueryHandler == nil {
		unregistered = append(unregistered, "entries.SearchLogQueryHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *RekorServerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *RekorServerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *RekorServerAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *RekorServerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "application/yaml":
			result["application/yaml"] = o.YamlConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *RekorServerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/x-pem-file":
			result["application/x-pem-file"] = o.ApplicationXPemFileProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "application/yaml":
			result["application/yaml"] = o.YamlProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *RekorServerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the rekor server API
func (o *RekorServerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *RekorServerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1/log/entries"] = entries.NewCreateLogEntry(o.context, o.EntriesCreateLogEntryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/log/entries"] = entries.NewGetLogEntryByIndex(o.context, o.EntriesGetLogEntryByIndexHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/log/entries/{entryUUID}"] = entries.NewGetLogEntryByUUID(o.context, o.EntriesGetLogEntryByUUIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/log/entries/{entryUUID}/proof"] = entries.NewGetLogEntryProof(o.context, o.EntriesGetLogEntryProofHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/log"] = tlog.NewGetLogInfo(o.context, o.TlogGetLogInfoHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/log/proof"] = tlog.NewGetLogProof(o.context, o.TlogGetLogProofHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1/log/publicKey"] = tlog.NewGetPublicKey(o.context, o.TlogGetPublicKeyHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1/log/entries/retrieve"] = entries.NewSearchLogQuery(o.context, o.EntriesSearchLogQueryHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *RekorServerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *RekorServerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *RekorServerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *RekorServerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *RekorServerAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
