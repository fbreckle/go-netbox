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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimVirtualDeviceContextsUpdateReader is a Reader for the DcimVirtualDeviceContextsUpdate structure.
type DcimVirtualDeviceContextsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualDeviceContextsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualDeviceContextsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualDeviceContextsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualDeviceContextsUpdateOK creates a DcimVirtualDeviceContextsUpdateOK with default headers values
func NewDcimVirtualDeviceContextsUpdateOK() *DcimVirtualDeviceContextsUpdateOK {
	return &DcimVirtualDeviceContextsUpdateOK{}
}

/*
DcimVirtualDeviceContextsUpdateOK describes a response with status code 200, with default header values.

DcimVirtualDeviceContextsUpdateOK dcim virtual device contexts update o k
*/
type DcimVirtualDeviceContextsUpdateOK struct {
	Payload *models.VirtualDeviceContext
}

// IsSuccess returns true when this dcim virtual device contexts update o k response has a 2xx status code
func (o *DcimVirtualDeviceContextsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual device contexts update o k response has a 3xx status code
func (o *DcimVirtualDeviceContextsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual device contexts update o k response has a 4xx status code
func (o *DcimVirtualDeviceContextsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual device contexts update o k response has a 5xx status code
func (o *DcimVirtualDeviceContextsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual device contexts update o k response a status code equal to that given
func (o *DcimVirtualDeviceContextsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimVirtualDeviceContextsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualDeviceContextsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualDeviceContextsUpdateOK) GetPayload() *models.VirtualDeviceContext {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDeviceContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualDeviceContextsUpdateDefault creates a DcimVirtualDeviceContextsUpdateDefault with default headers values
func NewDcimVirtualDeviceContextsUpdateDefault(code int) *DcimVirtualDeviceContextsUpdateDefault {
	return &DcimVirtualDeviceContextsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualDeviceContextsUpdateDefault describes a response with status code -1, with default header values.

DcimVirtualDeviceContextsUpdateDefault dcim virtual device contexts update default
*/
type DcimVirtualDeviceContextsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual device contexts update default response
func (o *DcimVirtualDeviceContextsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim virtual device contexts update default response has a 2xx status code
func (o *DcimVirtualDeviceContextsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual device contexts update default response has a 3xx status code
func (o *DcimVirtualDeviceContextsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual device contexts update default response has a 4xx status code
func (o *DcimVirtualDeviceContextsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual device contexts update default response has a 5xx status code
func (o *DcimVirtualDeviceContextsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual device contexts update default response a status code equal to that given
func (o *DcimVirtualDeviceContextsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimVirtualDeviceContextsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-device-contexts/{id}/][%d] dcim_virtual-device-contexts_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualDeviceContextsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/virtual-device-contexts/{id}/][%d] dcim_virtual-device-contexts_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualDeviceContextsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}