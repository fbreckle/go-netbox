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

// NewUsersOwnerGroupsListParams creates a new UsersOwnerGroupsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersOwnerGroupsListParams() *UsersOwnerGroupsListParams {
	return &UsersOwnerGroupsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersOwnerGroupsListParamsWithTimeout creates a new UsersOwnerGroupsListParams object
// with the ability to set a timeout on a request.
func NewUsersOwnerGroupsListParamsWithTimeout(timeout time.Duration) *UsersOwnerGroupsListParams {
	return &UsersOwnerGroupsListParams{
		timeout: timeout,
	}
}

// NewUsersOwnerGroupsListParamsWithContext creates a new UsersOwnerGroupsListParams object
// with the ability to set a context for a request.
func NewUsersOwnerGroupsListParamsWithContext(ctx context.Context) *UsersOwnerGroupsListParams {
	return &UsersOwnerGroupsListParams{
		Context: ctx,
	}
}

// NewUsersOwnerGroupsListParamsWithHTTPClient creates a new UsersOwnerGroupsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersOwnerGroupsListParamsWithHTTPClient(client *http.Client) *UsersOwnerGroupsListParams {
	return &UsersOwnerGroupsListParams{
		HTTPClient: client,
	}
}

/*
UsersOwnerGroupsListParams contains all the parameters to send to the API endpoint

	for the users owner groups list operation.

	Typically these are written to a http.Request.
*/
type UsersOwnerGroupsListParams struct {

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// Limit.
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	// Offset.
	Offset *int64

	// Ordering.
	Ordering *string

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users owner groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersOwnerGroupsListParams) WithDefaults() *UsersOwnerGroupsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users owner groups list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersOwnerGroupsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithTimeout(timeout time.Duration) *UsersOwnerGroupsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithContext(ctx context.Context) *UsersOwnerGroupsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithHTTPClient(client *http.Client) *UsersOwnerGroupsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the iD to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithID(iD *string) *UsersOwnerGroupsListParams {
	o.SetID(iD)
	return o
}

// SetID adds the iD to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetID(iD *string) {
	o.ID = iD
}

// WithIDGt adds the iDGt to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithIDGt(iDGt *string) *UsersOwnerGroupsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the iDGt to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithIDGte(iDGte *string) *UsersOwnerGroupsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the iDGte to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithIDLt(iDLt *string) *UsersOwnerGroupsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the iDLt to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithIDLte(iDLte *string) *UsersOwnerGroupsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the iDLte to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithIDn(iDn *string) *UsersOwnerGroupsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the iDn to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithLimit(limit *int64) *UsersOwnerGroupsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithName(name *string) *UsersOwnerGroupsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameEmpty(nameEmpty *string) *UsersOwnerGroupsListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameIc(nameIc *string) *UsersOwnerGroupsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameIe(nameIe *string) *UsersOwnerGroupsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameIew(nameIew *string) *UsersOwnerGroupsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameIsw(nameIsw *string) *UsersOwnerGroupsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNamen(namen *string) *UsersOwnerGroupsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the namen to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameNic(nameNic *string) *UsersOwnerGroupsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameNie(nameNie *string) *UsersOwnerGroupsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameNiew(nameNiew *string) *UsersOwnerGroupsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithNameNisw(nameNisw *string) *UsersOwnerGroupsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithOffset(offset *int64) *UsersOwnerGroupsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithOrdering(ordering *string) *UsersOwnerGroupsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithQ adds the q to the users owner groups list params
func (o *UsersOwnerGroupsListParams) WithQ(q *string) *UsersOwnerGroupsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the users owner groups list params
func (o *UsersOwnerGroupsListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *UsersOwnerGroupsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
