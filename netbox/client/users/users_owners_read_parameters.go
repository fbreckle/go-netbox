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
	"github.com/go-openapi/swag"
)

// NewUsersOwnersReadParams creates a new UsersOwnersReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersOwnersReadParams() *UsersOwnersReadParams {
	return &UsersOwnersReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersOwnersReadParamsWithTimeout creates a new UsersOwnersReadParams object
// with the ability to set a timeout on a request.
func NewUsersOwnersReadParamsWithTimeout(timeout time.Duration) *UsersOwnersReadParams {
	return &UsersOwnersReadParams{
		timeout: timeout,
	}
}

// NewUsersOwnersReadParamsWithContext creates a new UsersOwnersReadParams object
// with the ability to set a context for a request.
func NewUsersOwnersReadParamsWithContext(ctx context.Context) *UsersOwnersReadParams {
	return &UsersOwnersReadParams{
		Context: ctx,
	}
}

// NewUsersOwnersReadParamsWithHTTPClient creates a new UsersOwnersReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersOwnersReadParamsWithHTTPClient(client *http.Client) *UsersOwnersReadParams {
	return &UsersOwnersReadParams{
		HTTPClient: client,
	}
}

/*
UsersOwnersReadParams contains all the parameters to send to the API endpoint

	for the users owners read operation.

	Typically these are written to a http.Request.
*/
type UsersOwnersReadParams struct {

	/* ID.

	   A unique integer value identifying this owner.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users owners read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersOwnersReadParams) WithDefaults() *UsersOwnersReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users owners read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersOwnersReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users owners read params
func (o *UsersOwnersReadParams) WithTimeout(timeout time.Duration) *UsersOwnersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users owners read params
func (o *UsersOwnersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users owners read params
func (o *UsersOwnersReadParams) WithContext(ctx context.Context) *UsersOwnersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users owners read params
func (o *UsersOwnersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users owners read params
func (o *UsersOwnersReadParams) WithHTTPClient(client *http.Client) *UsersOwnersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users owners read params
func (o *UsersOwnersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users owners read params
func (o *UsersOwnersReadParams) WithID(id int64) *UsersOwnersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users owners read params
func (o *UsersOwnersReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersOwnersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
