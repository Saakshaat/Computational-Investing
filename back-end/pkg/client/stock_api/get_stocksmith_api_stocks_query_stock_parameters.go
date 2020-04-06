// Code generated by go-swagger; DO NOT EDIT.

package stock_api

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
)

// NewGetStocksmithAPIStocksQueryStockParams creates a new GetStocksmithAPIStocksQueryStockParams object
// with the default values initialized.
func NewGetStocksmithAPIStocksQueryStockParams() *GetStocksmithAPIStocksQueryStockParams {

	return &GetStocksmithAPIStocksQueryStockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStocksmithAPIStocksQueryStockParamsWithTimeout creates a new GetStocksmithAPIStocksQueryStockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStocksmithAPIStocksQueryStockParamsWithTimeout(timeout time.Duration) *GetStocksmithAPIStocksQueryStockParams {

	return &GetStocksmithAPIStocksQueryStockParams{

		timeout: timeout,
	}
}

// NewGetStocksmithAPIStocksQueryStockParamsWithContext creates a new GetStocksmithAPIStocksQueryStockParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStocksmithAPIStocksQueryStockParamsWithContext(ctx context.Context) *GetStocksmithAPIStocksQueryStockParams {

	return &GetStocksmithAPIStocksQueryStockParams{

		Context: ctx,
	}
}

// NewGetStocksmithAPIStocksQueryStockParamsWithHTTPClient creates a new GetStocksmithAPIStocksQueryStockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStocksmithAPIStocksQueryStockParamsWithHTTPClient(client *http.Client) *GetStocksmithAPIStocksQueryStockParams {

	return &GetStocksmithAPIStocksQueryStockParams{
		HTTPClient: client,
	}
}

/*GetStocksmithAPIStocksQueryStockParams contains all the parameters to send to the API endpoint
for the get stocksmith API stocks query stock operation typically these are written to a http.Request
*/
type GetStocksmithAPIStocksQueryStockParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get stocksmith API stocks query stock params
func (o *GetStocksmithAPIStocksQueryStockParams) WithTimeout(timeout time.Duration) *GetStocksmithAPIStocksQueryStockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get stocksmith API stocks query stock params
func (o *GetStocksmithAPIStocksQueryStockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get stocksmith API stocks query stock params
func (o *GetStocksmithAPIStocksQueryStockParams) WithContext(ctx context.Context) *GetStocksmithAPIStocksQueryStockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get stocksmith API stocks query stock params
func (o *GetStocksmithAPIStocksQueryStockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get stocksmith API stocks query stock params
func (o *GetStocksmithAPIStocksQueryStockParams) WithHTTPClient(client *http.Client) *GetStocksmithAPIStocksQueryStockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get stocksmith API stocks query stock params
func (o *GetStocksmithAPIStocksQueryStockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStocksmithAPIStocksQueryStockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
