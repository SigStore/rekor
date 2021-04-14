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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JarV001Schema JAR v0.0.1 Schema
//
// Schema for JAR entries
//
// swagger:model jarV001Schema
type JarV001Schema struct {

	// archive
	// Required: true
	Archive *JarV001SchemaArchive `json:"archive"`

	// Arbitrary content to be included in the verifiable entry in the transparency log
	ExtraData interface{} `json:"extraData,omitempty"`

	// signature
	Signature *JarV001SchemaSignature `json:"signature,omitempty"`
}

// Validate validates this jar v001 schema
func (m *JarV001Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001Schema) validateArchive(formats strfmt.Registry) error {

	if err := validate.Required("archive", "body", m.Archive); err != nil {
		return err
	}

	if m.Archive != nil {
		if err := m.Archive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive")
			}
			return err
		}
	}

	return nil
}

func (m *JarV001Schema) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(m.Signature) { // not required
		return nil
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jar v001 schema based on the context it is used
func (m *JarV001Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchive(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001Schema) contextValidateArchive(ctx context.Context, formats strfmt.Registry) error {

	if m.Archive != nil {
		if err := m.Archive.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive")
			}
			return err
		}
	}

	return nil
}

func (m *JarV001Schema) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {
		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JarV001Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JarV001Schema) UnmarshalBinary(b []byte) error {
	var res JarV001Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JarV001SchemaArchive Information about the archive associated with the entry
//
// swagger:model JarV001SchemaArchive
type JarV001SchemaArchive struct {

	// Specifies the archive inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// hash
	Hash *JarV001SchemaArchiveHash `json:"hash,omitempty"`

	// Specifies the location of the archive; if this is specified, a hash value must also be provided
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this jar v001 schema archive
func (m *JarV001SchemaArchive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001SchemaArchive) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *JarV001SchemaArchive) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("archive"+"."+"url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this jar v001 schema archive based on the context it is used
func (m *JarV001SchemaArchive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001SchemaArchive) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {
		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archive" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JarV001SchemaArchive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JarV001SchemaArchive) UnmarshalBinary(b []byte) error {
	var res JarV001SchemaArchive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JarV001SchemaArchiveHash Specifies the hash algorithm and value encompassing the entire signed archive
//
// swagger:model JarV001SchemaArchiveHash
type JarV001SchemaArchiveHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value for the archive
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this jar v001 schema archive hash
func (m *JarV001SchemaArchiveHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jarV001SchemaArchiveHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jarV001SchemaArchiveHashTypeAlgorithmPropEnum = append(jarV001SchemaArchiveHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// JarV001SchemaArchiveHashAlgorithmSha256 captures enum value "sha256"
	JarV001SchemaArchiveHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *JarV001SchemaArchiveHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jarV001SchemaArchiveHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JarV001SchemaArchiveHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("archive"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("archive"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *JarV001SchemaArchiveHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("archive"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this jar v001 schema archive hash based on context it is used
func (m *JarV001SchemaArchiveHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JarV001SchemaArchiveHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JarV001SchemaArchiveHash) UnmarshalBinary(b []byte) error {
	var res JarV001SchemaArchiveHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JarV001SchemaSignature Information about the included signature in the JAR file
//
// swagger:model JarV001SchemaSignature
type JarV001SchemaSignature struct {

	// Specifies the PKCS7 signature embedded within the JAR file
	// Required: true
	// Format: byte
	Content *strfmt.Base64 `json:"content"`

	// public key
	// Required: true
	PublicKey *JarV001SchemaSignaturePublicKey `json:"publicKey"`
}

// Validate validates this jar v001 schema signature
func (m *JarV001SchemaSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001SchemaSignature) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("signature"+"."+"content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *JarV001SchemaSignature) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("signature"+"."+"publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jar v001 schema signature based on the context it is used
func (m *JarV001SchemaSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001SchemaSignature) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicKey != nil {
		if err := m.PublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JarV001SchemaSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JarV001SchemaSignature) UnmarshalBinary(b []byte) error {
	var res JarV001SchemaSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JarV001SchemaSignaturePublicKey The X509 certificate containing the public key JAR which verifies the signature of the JAR
//
// swagger:model JarV001SchemaSignaturePublicKey
type JarV001SchemaSignaturePublicKey struct {

	// Specifies the content of the X509 certificate containing the public key used to verify the signature
	// Required: true
	// Format: byte
	Content *strfmt.Base64 `json:"content"`
}

// Validate validates this jar v001 schema signature public key
func (m *JarV001SchemaSignaturePublicKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JarV001SchemaSignaturePublicKey) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("signature"+"."+"publicKey"+"."+"content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this jar v001 schema signature public key based on context it is used
func (m *JarV001SchemaSignaturePublicKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JarV001SchemaSignaturePublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JarV001SchemaSignaturePublicKey) UnmarshalBinary(b []byte) error {
	var res JarV001SchemaSignaturePublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
