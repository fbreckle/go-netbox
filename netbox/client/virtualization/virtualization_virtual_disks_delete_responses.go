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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationVirtualDisksDeleteReader is a Reader for the VirtualizationVirtualDisksDelete structure.
type VirtualizationVirtualDisksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualDisksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationVirtualDisksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationVirtualDisksDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualDisksDeleteNoContent creates a VirtualizationVirtualDisksDeleteNoContent with default headers values
func NewVirtualizationVirtualDisksDeleteNoContent() *VirtualizationVirtualDisksDeleteNoContent {
	return &VirtualizationVirtualDisksDeleteNoContent{}
}

/*
VirtualizationVirtualDisksDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationVirtualDisksDeleteNoContent virtualization virtual disks delete no content
*/
type VirtualizationVirtualDisksDeleteNoContent struct {
}

// IsSuccess returns true when this virtualization virtual disks delete no content response has a 2xx status code
func (o *VirtualizationVirtualDisksDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual disks delete no content response has a 3xx status code
func (o *VirtualizationVirtualDisksDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual disks delete no content response has a 4xx status code
func (o *VirtualizationVirtualDisksDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual disks delete no content response has a 5xx status code
func (o *VirtualizationVirtualDisksDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual disks delete no content response a status code equal to that given
func (o *VirtualizationVirtualDisksDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the virtualization virtual disks delete no content response
func (o *VirtualizationVirtualDisksDeleteNoContent) Code() int {
	return 204
}

func (o *VirtualizationVirtualDisksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-disks/{id}/][%d] virtualizationVirtualDisksDeleteNoContent", 204)
}

func (o *VirtualizationVirtualDisksDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtualization/virtual-disks/{id}/][%d] virtualizationVirtualDisksDeleteNoContent", 204)
}

func (o *VirtualizationVirtualDisksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationVirtualDisksDeleteDefault creates a VirtualizationVirtualDisksDeleteDefault with default headers values
func NewVirtualizationVirtualDisksDeleteDefault(code int) *VirtualizationVirtualDisksDeleteDefault {
	return &VirtualizationVirtualDisksDeleteDefault{
		_statusCode: code,
	}
}

/*
VirtualizationVirtualDisksDeleteDefault describes a response with status code -1, with default header values.

VirtualizationVirtualDisksDeleteDefault virtualization virtual disks delete default
*/
type VirtualizationVirtualDisksDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this virtualization virtual disks delete default response has a 2xx status code
func (o *VirtualizationVirtualDisksDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization virtual disks delete default response has a 3xx status code
func (o *VirtualizationVirtualDisksDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization virtual disks delete default response has a 4xx status code
func (o *VirtualizationVirtualDisksDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization virtual disks delete default response has a 5xx status code
func (o *VirtualizationVirtualDisksDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization virtual disks delete default response a status code equal to that given
func (o *VirtualizationVirtualDisksDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the virtualization virtual disks delete default response
func (o *VirtualizationVirtualDisksDeleteDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationVirtualDisksDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /virtualization/virtual-disks/{id}/][%d] virtualization_virtual-disks_delete default %s", o._statusCode, payload)
}

func (o *VirtualizationVirtualDisksDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /virtualization/virtual-disks/{id}/][%d] virtualization_virtual-disks_delete default %s", o._statusCode, payload)
}

func (o *VirtualizationVirtualDisksDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualDisksDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
