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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// ExtrasSavedFiltersUpdateReader is a Reader for the ExtrasSavedFiltersUpdate structure.
type ExtrasSavedFiltersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasSavedFiltersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasSavedFiltersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasSavedFiltersUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasSavedFiltersUpdateOK creates a ExtrasSavedFiltersUpdateOK with default headers values
func NewExtrasSavedFiltersUpdateOK() *ExtrasSavedFiltersUpdateOK {
	return &ExtrasSavedFiltersUpdateOK{}
}

/*
ExtrasSavedFiltersUpdateOK describes a response with status code 200, with default header values.

ExtrasSavedFiltersUpdateOK extras saved filters update o k
*/
type ExtrasSavedFiltersUpdateOK struct {
	Payload *models.SavedFilter
}

// IsSuccess returns true when this extras saved filters update o k response has a 2xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras saved filters update o k response has a 3xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras saved filters update o k response has a 4xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras saved filters update o k response has a 5xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras saved filters update o k response a status code equal to that given
func (o *ExtrasSavedFiltersUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasSavedFiltersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extrasSavedFiltersUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extrasSavedFiltersUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateOK) GetPayload() *models.SavedFilter {
	return o.Payload
}

func (o *ExtrasSavedFiltersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasSavedFiltersUpdateDefault creates a ExtrasSavedFiltersUpdateDefault with default headers values
func NewExtrasSavedFiltersUpdateDefault(code int) *ExtrasSavedFiltersUpdateDefault {
	return &ExtrasSavedFiltersUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasSavedFiltersUpdateDefault describes a response with status code -1, with default header values.

ExtrasSavedFiltersUpdateDefault extras saved filters update default
*/
type ExtrasSavedFiltersUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras saved filters update default response
func (o *ExtrasSavedFiltersUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras saved filters update default response has a 2xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras saved filters update default response has a 3xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras saved filters update default response has a 4xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras saved filters update default response has a 5xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras saved filters update default response a status code equal to that given
func (o *ExtrasSavedFiltersUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasSavedFiltersUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extras_saved-filters_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extras_saved-filters_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasSavedFiltersUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}