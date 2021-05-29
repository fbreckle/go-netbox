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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasContentTypesReadParams creates a new ExtrasContentTypesReadParams object
// with the default values initialized.
func NewExtrasContentTypesReadParams() *ExtrasContentTypesReadParams {
	var ()
	return &ExtrasContentTypesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasContentTypesReadParamsWithTimeout creates a new ExtrasContentTypesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasContentTypesReadParamsWithTimeout(timeout time.Duration) *ExtrasContentTypesReadParams {
	var ()
	return &ExtrasContentTypesReadParams{

		timeout: timeout,
	}
}

// NewExtrasContentTypesReadParamsWithContext creates a new ExtrasContentTypesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasContentTypesReadParamsWithContext(ctx context.Context) *ExtrasContentTypesReadParams {
	var ()
	return &ExtrasContentTypesReadParams{

		Context: ctx,
	}
}

// NewExtrasContentTypesReadParamsWithHTTPClient creates a new ExtrasContentTypesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasContentTypesReadParamsWithHTTPClient(client *http.Client) *ExtrasContentTypesReadParams {
	var ()
	return &ExtrasContentTypesReadParams{
		HTTPClient: client,
	}
}

/*ExtrasContentTypesReadParams contains all the parameters to send to the API endpoint
for the extras content types read operation typically these are written to a http.Request
*/
type ExtrasContentTypesReadParams struct {

	/*ID
	  A unique integer value identifying this content type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras content types read params
func (o *ExtrasContentTypesReadParams) WithTimeout(timeout time.Duration) *ExtrasContentTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras content types read params
func (o *ExtrasContentTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras content types read params
func (o *ExtrasContentTypesReadParams) WithContext(ctx context.Context) *ExtrasContentTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras content types read params
func (o *ExtrasContentTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras content types read params
func (o *ExtrasContentTypesReadParams) WithHTTPClient(client *http.Client) *ExtrasContentTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras content types read params
func (o *ExtrasContentTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras content types read params
func (o *ExtrasContentTypesReadParams) WithID(id int64) *ExtrasContentTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras content types read params
func (o *ExtrasContentTypesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasContentTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}