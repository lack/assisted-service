// Code generated by go-swagger; DO NOT EDIT.

package installer

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

	"github.com/openshift/assisted-service/models"
)

// NewPostStepReplyParams creates a new PostStepReplyParams object
// with the default values initialized.
func NewPostStepReplyParams() *PostStepReplyParams {
	var ()
	return &PostStepReplyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStepReplyParamsWithTimeout creates a new PostStepReplyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStepReplyParamsWithTimeout(timeout time.Duration) *PostStepReplyParams {
	var ()
	return &PostStepReplyParams{

		timeout: timeout,
	}
}

// NewPostStepReplyParamsWithContext creates a new PostStepReplyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostStepReplyParamsWithContext(ctx context.Context) *PostStepReplyParams {
	var ()
	return &PostStepReplyParams{

		Context: ctx,
	}
}

// NewPostStepReplyParamsWithHTTPClient creates a new PostStepReplyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostStepReplyParamsWithHTTPClient(client *http.Client) *PostStepReplyParams {
	var ()
	return &PostStepReplyParams{
		HTTPClient: client,
	}
}

/*PostStepReplyParams contains all the parameters to send to the API endpoint
for the post step reply operation typically these are written to a http.Request
*/
type PostStepReplyParams struct {

	/*ClusterID*/
	ClusterID strfmt.UUID
	/*DiscoveryAgentVersion*/
	DiscoveryAgentVersion *string
	/*HostID*/
	HostID strfmt.UUID
	/*Reply*/
	Reply *models.StepReply

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post step reply params
func (o *PostStepReplyParams) WithTimeout(timeout time.Duration) *PostStepReplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post step reply params
func (o *PostStepReplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post step reply params
func (o *PostStepReplyParams) WithContext(ctx context.Context) *PostStepReplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post step reply params
func (o *PostStepReplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post step reply params
func (o *PostStepReplyParams) WithHTTPClient(client *http.Client) *PostStepReplyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post step reply params
func (o *PostStepReplyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the post step reply params
func (o *PostStepReplyParams) WithClusterID(clusterID strfmt.UUID) *PostStepReplyParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the post step reply params
func (o *PostStepReplyParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithDiscoveryAgentVersion adds the discoveryAgentVersion to the post step reply params
func (o *PostStepReplyParams) WithDiscoveryAgentVersion(discoveryAgentVersion *string) *PostStepReplyParams {
	o.SetDiscoveryAgentVersion(discoveryAgentVersion)
	return o
}

// SetDiscoveryAgentVersion adds the discoveryAgentVersion to the post step reply params
func (o *PostStepReplyParams) SetDiscoveryAgentVersion(discoveryAgentVersion *string) {
	o.DiscoveryAgentVersion = discoveryAgentVersion
}

// WithHostID adds the hostID to the post step reply params
func (o *PostStepReplyParams) WithHostID(hostID strfmt.UUID) *PostStepReplyParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the post step reply params
func (o *PostStepReplyParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithReply adds the reply to the post step reply params
func (o *PostStepReplyParams) WithReply(reply *models.StepReply) *PostStepReplyParams {
	o.SetReply(reply)
	return o
}

// SetReply adds the reply to the post step reply params
func (o *PostStepReplyParams) SetReply(reply *models.StepReply) {
	o.Reply = reply
}

// WriteToRequest writes these params to a swagger request
func (o *PostStepReplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.DiscoveryAgentVersion != nil {

		// header param discovery_agent_version
		if err := r.SetHeaderParam("discovery_agent_version", *o.DiscoveryAgentVersion); err != nil {
			return err
		}

	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if o.Reply != nil {
		if err := r.SetBodyParam(o.Reply); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
