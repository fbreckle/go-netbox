// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SecretsSecretsBulkDeleteReader is a Reader for the SecretsSecretsBulkDelete structure.
type SecretsSecretsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSecretsSecretsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretsBulkDeleteNoContent creates a SecretsSecretsBulkDeleteNoContent with default headers values
func NewSecretsSecretsBulkDeleteNoContent() *SecretsSecretsBulkDeleteNoContent {
	return &SecretsSecretsBulkDeleteNoContent{}
}

/*SecretsSecretsBulkDeleteNoContent handles this case with default header values.

SecretsSecretsBulkDeleteNoContent secrets secrets bulk delete no content
*/
type SecretsSecretsBulkDeleteNoContent struct {
}

func (o *SecretsSecretsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /secrets/secrets/][%d] secretsSecretsBulkDeleteNoContent ", 204)
}

func (o *SecretsSecretsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSecretsSecretsBulkDeleteDefault creates a SecretsSecretsBulkDeleteDefault with default headers values
func NewSecretsSecretsBulkDeleteDefault(code int) *SecretsSecretsBulkDeleteDefault {
	return &SecretsSecretsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*SecretsSecretsBulkDeleteDefault handles this case with default header values.

SecretsSecretsBulkDeleteDefault secrets secrets bulk delete default
*/
type SecretsSecretsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secrets bulk delete default response
func (o *SecretsSecretsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /secrets/secrets/][%d] secrets_secrets_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
