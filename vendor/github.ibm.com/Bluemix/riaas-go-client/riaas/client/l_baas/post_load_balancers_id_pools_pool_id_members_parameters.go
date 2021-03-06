// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewPostLoadBalancersIDPoolsPoolIDMembersParams creates a new PostLoadBalancersIDPoolsPoolIDMembersParams object
// with the default values initialized.
func NewPostLoadBalancersIDPoolsPoolIDMembersParams() *PostLoadBalancersIDPoolsPoolIDMembersParams {
	var ()
	return &PostLoadBalancersIDPoolsPoolIDMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithTimeout creates a new PostLoadBalancersIDPoolsPoolIDMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithTimeout(timeout time.Duration) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	var ()
	return &PostLoadBalancersIDPoolsPoolIDMembersParams{

		timeout: timeout,
	}
}

// NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithContext creates a new PostLoadBalancersIDPoolsPoolIDMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithContext(ctx context.Context) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	var ()
	return &PostLoadBalancersIDPoolsPoolIDMembersParams{

		Context: ctx,
	}
}

// NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithHTTPClient creates a new PostLoadBalancersIDPoolsPoolIDMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLoadBalancersIDPoolsPoolIDMembersParamsWithHTTPClient(client *http.Client) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	var ()
	return &PostLoadBalancersIDPoolsPoolIDMembersParams{
		HTTPClient: client,
	}
}

/*PostLoadBalancersIDPoolsPoolIDMembersParams contains all the parameters to send to the API endpoint
for the post load balancers ID pools pool ID members operation typically these are written to a http.Request
*/
type PostLoadBalancersIDPoolsPoolIDMembersParams struct {

	/*Body
	  The member template

	*/
	Body *models.MemberTemplate
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The load balancer identifier

	*/
	ID string
	/*PoolID
	  The pool identifier

	*/
	PoolID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithTimeout(timeout time.Duration) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithContext(ctx context.Context) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithHTTPClient(client *http.Client) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithBody(body *models.MemberTemplate) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetBody(body *models.MemberTemplate) {
	o.Body = body
}

// WithGeneration adds the generation to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithGeneration(generation int64) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithID(id string) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetID(id string) {
	o.ID = id
}

// WithPoolID adds the poolID to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithPoolID(poolID string) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetPoolID(poolID)
	return o
}

// SetPoolID adds the poolId to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetPoolID(poolID string) {
	o.PoolID = poolID
}

// WithVersion adds the version to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WithVersion(version string) *PostLoadBalancersIDPoolsPoolIDMembersParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post load balancers ID pools pool ID members params
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostLoadBalancersIDPoolsPoolIDMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param pool_id
	if err := r.SetPathParam("pool_id", o.PoolID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
