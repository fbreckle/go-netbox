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

package users

// This file was manually added to extend the client with NetBox's Owner/OwnerGroup
// endpoints (netbox-community/netbox users app), which are not yet covered by the
// upstream swagger-generated client.

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

// NewUsersOwnerGroupsCreateParams creates a new UsersOwnerGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersOwnerGroupsCreateParams() *UsersOwnerGroupsCreateParams {
	return &UsersOwnerGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersOwnerGroupsCreateParamsWithTimeout creates a new UsersOwnerGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewUsersOwnerGroupsCreateParamsWithTimeout(timeout time.Duration) *UsersOwnerGroupsCreateParams {
	return &UsersOwnerGroupsCreateParams{
		timeout: timeout,
	}
}

// NewUsersOwnerGroupsCreateParamsWithContext creates a new UsersOwnerGroupsCreateParams object
// with the ability to set a context for a request.
func NewUsersOwnerGroupsCreateParamsWithContext(ctx context.Context) *UsersOwnerGroupsCreateParams {
	return &UsersOwnerGroupsCreateParams{
		Context: ctx,
	}
}

// NewUsersOwnerGroupsCreateParamsWithHTTPClient creates a new UsersOwnerGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersOwnerGroupsCreateParamsWithHTTPClient(client *http.Client) *UsersOwnerGroupsCreateParams {
	return &UsersOwnerGroupsCreateParams{
		HTTPClient: client,
	}
}

/*
UsersOwnerGroupsCreateParams contains all the parameters to send to the API endpoint

	for the users owner groups create operation.

	Typically these are written to a http.Request.
*/
type UsersOwnerGroupsCreateParams struct {

	// Data.
	Data *models.OwnerGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users owner groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersOwnerGroupsCreateParams) WithDefaults() *UsersOwnerGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users owner groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersOwnerGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) WithTimeout(timeout time.Duration) *UsersOwnerGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) WithContext(ctx context.Context) *UsersOwnerGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) WithHTTPClient(client *http.Client) *UsersOwnerGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) WithData(data *models.OwnerGroup) *UsersOwnerGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users owner groups create params
func (o *UsersOwnerGroupsCreateParams) SetData(data *models.OwnerGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersOwnerGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
