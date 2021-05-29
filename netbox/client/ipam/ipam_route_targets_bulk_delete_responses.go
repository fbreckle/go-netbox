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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRouteTargetsBulkDeleteReader is a Reader for the IpamRouteTargetsBulkDelete structure.
type IpamRouteTargetsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRouteTargetsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRouteTargetsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsBulkDeleteNoContent creates a IpamRouteTargetsBulkDeleteNoContent with default headers values
func NewIpamRouteTargetsBulkDeleteNoContent() *IpamRouteTargetsBulkDeleteNoContent {
	return &IpamRouteTargetsBulkDeleteNoContent{}
}

/*IpamRouteTargetsBulkDeleteNoContent handles this case with default header values.

IpamRouteTargetsBulkDeleteNoContent ipam route targets bulk delete no content
*/
type IpamRouteTargetsBulkDeleteNoContent struct {
}

func (o *IpamRouteTargetsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipamRouteTargetsBulkDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRouteTargetsBulkDeleteDefault creates a IpamRouteTargetsBulkDeleteDefault with default headers values
func NewIpamRouteTargetsBulkDeleteDefault(code int) *IpamRouteTargetsBulkDeleteDefault {
	return &IpamRouteTargetsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*IpamRouteTargetsBulkDeleteDefault handles this case with default header values.

IpamRouteTargetsBulkDeleteDefault ipam route targets bulk delete default
*/
type IpamRouteTargetsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets bulk delete default response
func (o *IpamRouteTargetsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamRouteTargetsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/][%d] ipam_route-targets_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}