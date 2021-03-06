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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// SecretsSecretsBulkPartialUpdateReader is a Reader for the SecretsSecretsBulkPartialUpdate structure.
type SecretsSecretsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretsSecretsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecretsSecretsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecretsSecretsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecretsSecretsBulkPartialUpdateOK creates a SecretsSecretsBulkPartialUpdateOK with default headers values
func NewSecretsSecretsBulkPartialUpdateOK() *SecretsSecretsBulkPartialUpdateOK {
	return &SecretsSecretsBulkPartialUpdateOK{}
}

/*SecretsSecretsBulkPartialUpdateOK handles this case with default header values.

SecretsSecretsBulkPartialUpdateOK secrets secrets bulk partial update o k
*/
type SecretsSecretsBulkPartialUpdateOK struct {
	Payload *models.Secret
}

func (o *SecretsSecretsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /secrets/secrets/][%d] secretsSecretsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *SecretsSecretsBulkPartialUpdateOK) GetPayload() *models.Secret {
	return o.Payload
}

func (o *SecretsSecretsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretsSecretsBulkPartialUpdateDefault creates a SecretsSecretsBulkPartialUpdateDefault with default headers values
func NewSecretsSecretsBulkPartialUpdateDefault(code int) *SecretsSecretsBulkPartialUpdateDefault {
	return &SecretsSecretsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*SecretsSecretsBulkPartialUpdateDefault handles this case with default header values.

SecretsSecretsBulkPartialUpdateDefault secrets secrets bulk partial update default
*/
type SecretsSecretsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the secrets secrets bulk partial update default response
func (o *SecretsSecretsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *SecretsSecretsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /secrets/secrets/][%d] secrets_secrets_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *SecretsSecretsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *SecretsSecretsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
