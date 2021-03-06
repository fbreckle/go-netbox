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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// NewDcimConsoleServerPortsBulkPartialUpdateParams creates a new DcimConsoleServerPortsBulkPartialUpdateParams object
// with the default values initialized.
func NewDcimConsoleServerPortsBulkPartialUpdateParams() *DcimConsoleServerPortsBulkPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortsBulkPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsBulkPartialUpdateParamsWithTimeout creates a new DcimConsoleServerPortsBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsoleServerPortsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsBulkPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortsBulkPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsBulkPartialUpdateParamsWithContext creates a new DcimConsoleServerPortsBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsoleServerPortsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortsBulkPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortsBulkPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimConsoleServerPortsBulkPartialUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortsBulkPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsoleServerPortsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsBulkPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimConsoleServerPortsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim console server ports bulk partial update operation typically these are written to a http.Request
*/
type DcimConsoleServerPortsBulkPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableConsoleServerPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) WithData(data *models.WritableConsoleServerPort) *DcimConsoleServerPortsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server ports bulk partial update params
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) SetData(data *models.WritableConsoleServerPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
