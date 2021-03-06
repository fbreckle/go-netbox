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

package circuits

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

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewCircuitsCircuitTerminationsBulkUpdateParams creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the default values initialized.
func NewCircuitsCircuitTerminationsBulkUpdateParams() *CircuitsCircuitTerminationsBulkUpdateParams {
	var ()
	return &CircuitsCircuitTerminationsBulkUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateParamsWithTimeout creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTerminationsBulkUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsBulkUpdateParams {
	var ()
	return &CircuitsCircuitTerminationsBulkUpdateParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateParamsWithContext creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTerminationsBulkUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsBulkUpdateParams {
	var ()
	return &CircuitsCircuitTerminationsBulkUpdateParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateParamsWithHTTPClient creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTerminationsBulkUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsBulkUpdateParams {
	var ()
	return &CircuitsCircuitTerminationsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTerminationsBulkUpdateParams contains all the parameters to send to the API endpoint
for the circuits circuit terminations bulk update operation typically these are written to a http.Request
*/
type CircuitsCircuitTerminationsBulkUpdateParams struct {

	/*Data*/
	Data *models.WritableCircuitTermination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithData(data *models.WritableCircuitTermination) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetData(data *models.WritableCircuitTermination) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
