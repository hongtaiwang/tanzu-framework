// Code generated by go-swagger; DO NOT EDIT.

package docker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// NewImportTKGConfigForDockerParams creates a new ImportTKGConfigForDockerParams object
// with the default values initialized.
func NewImportTKGConfigForDockerParams() *ImportTKGConfigForDockerParams {
	var ()
	return &ImportTKGConfigForDockerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportTKGConfigForDockerParamsWithTimeout creates a new ImportTKGConfigForDockerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportTKGConfigForDockerParamsWithTimeout(timeout time.Duration) *ImportTKGConfigForDockerParams {
	var ()
	return &ImportTKGConfigForDockerParams{

		timeout: timeout,
	}
}

// NewImportTKGConfigForDockerParamsWithContext creates a new ImportTKGConfigForDockerParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportTKGConfigForDockerParamsWithContext(ctx context.Context) *ImportTKGConfigForDockerParams {
	var ()
	return &ImportTKGConfigForDockerParams{

		Context: ctx,
	}
}

// NewImportTKGConfigForDockerParamsWithHTTPClient creates a new ImportTKGConfigForDockerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportTKGConfigForDockerParamsWithHTTPClient(client *http.Client) *ImportTKGConfigForDockerParams {
	var ()
	return &ImportTKGConfigForDockerParams{
		HTTPClient: client,
	}
}

/*
ImportTKGConfigForDockerParams contains all the parameters to send to the API endpoint
for the import t k g config for docker operation typically these are written to a http.Request
*/
type ImportTKGConfigForDockerParams struct {

	/*Params
	  config file from which to generate tkg configuration for docker

	*/
	Params *models.ConfigFile

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) WithTimeout(timeout time.Duration) *ImportTKGConfigForDockerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) WithContext(ctx context.Context) *ImportTKGConfigForDockerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) WithHTTPClient(client *http.Client) *ImportTKGConfigForDockerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParams adds the params to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) WithParams(params *models.ConfigFile) *ImportTKGConfigForDockerParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the import t k g config for docker params
func (o *ImportTKGConfigForDockerParams) SetParams(params *models.ConfigFile) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *ImportTKGConfigForDockerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Params != nil {
		if err := r.SetBodyParam(o.Params); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
