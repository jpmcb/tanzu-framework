// Code generated by go-swagger; DO NOT EDIT.

package aws

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

	models "github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/web/server/models"
)

// NewCreateAWSRegionalClusterParams creates a new CreateAWSRegionalClusterParams object
// with the default values initialized.
func NewCreateAWSRegionalClusterParams() *CreateAWSRegionalClusterParams {
	var ()
	return &CreateAWSRegionalClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAWSRegionalClusterParamsWithTimeout creates a new CreateAWSRegionalClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAWSRegionalClusterParamsWithTimeout(timeout time.Duration) *CreateAWSRegionalClusterParams {
	var ()
	return &CreateAWSRegionalClusterParams{

		timeout: timeout,
	}
}

// NewCreateAWSRegionalClusterParamsWithContext creates a new CreateAWSRegionalClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAWSRegionalClusterParamsWithContext(ctx context.Context) *CreateAWSRegionalClusterParams {
	var ()
	return &CreateAWSRegionalClusterParams{

		Context: ctx,
	}
}

// NewCreateAWSRegionalClusterParamsWithHTTPClient creates a new CreateAWSRegionalClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAWSRegionalClusterParamsWithHTTPClient(client *http.Client) *CreateAWSRegionalClusterParams {
	var ()
	return &CreateAWSRegionalClusterParams{
		HTTPClient: client,
	}
}

/*CreateAWSRegionalClusterParams contains all the parameters to send to the API endpoint
for the create a w s regional cluster operation typically these are written to a http.Request
*/
type CreateAWSRegionalClusterParams struct {

	/*Params
	  parameters to create a regional cluster

	*/
	Params *models.AWSRegionalClusterParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) WithTimeout(timeout time.Duration) *CreateAWSRegionalClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) WithContext(ctx context.Context) *CreateAWSRegionalClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) WithHTTPClient(client *http.Client) *CreateAWSRegionalClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParams adds the params to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) WithParams(params *models.AWSRegionalClusterParams) *CreateAWSRegionalClusterParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the create a w s regional cluster params
func (o *CreateAWSRegionalClusterParams) SetParams(params *models.AWSRegionalClusterParams) {
	o.Params = params
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAWSRegionalClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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