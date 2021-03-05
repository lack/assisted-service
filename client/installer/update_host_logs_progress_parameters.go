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

// NewUpdateHostLogsProgressParams creates a new UpdateHostLogsProgressParams object
// with the default values initialized.
func NewUpdateHostLogsProgressParams() *UpdateHostLogsProgressParams {
	var ()
	return &UpdateHostLogsProgressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateHostLogsProgressParamsWithTimeout creates a new UpdateHostLogsProgressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateHostLogsProgressParamsWithTimeout(timeout time.Duration) *UpdateHostLogsProgressParams {
	var ()
	return &UpdateHostLogsProgressParams{

		timeout: timeout,
	}
}

// NewUpdateHostLogsProgressParamsWithContext creates a new UpdateHostLogsProgressParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateHostLogsProgressParamsWithContext(ctx context.Context) *UpdateHostLogsProgressParams {
	var ()
	return &UpdateHostLogsProgressParams{

		Context: ctx,
	}
}

// NewUpdateHostLogsProgressParamsWithHTTPClient creates a new UpdateHostLogsProgressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateHostLogsProgressParamsWithHTTPClient(client *http.Client) *UpdateHostLogsProgressParams {
	var ()
	return &UpdateHostLogsProgressParams{
		HTTPClient: client,
	}
}

/*UpdateHostLogsProgressParams contains all the parameters to send to the API endpoint
for the update host logs progress operation typically these are written to a http.Request
*/
type UpdateHostLogsProgressParams struct {

	/*ClusterID
	  The cluster whose log progress is being updated.

	*/
	ClusterID strfmt.UUID
	/*HostID
	  The host whose log progress is being updated.

	*/
	HostID strfmt.UUID
	/*LogsProgressParams
	  Parameters for updating log progress.

	*/
	LogsProgressParams *models.LogsProgressParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update host logs progress params
func (o *UpdateHostLogsProgressParams) WithTimeout(timeout time.Duration) *UpdateHostLogsProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update host logs progress params
func (o *UpdateHostLogsProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update host logs progress params
func (o *UpdateHostLogsProgressParams) WithContext(ctx context.Context) *UpdateHostLogsProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update host logs progress params
func (o *UpdateHostLogsProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update host logs progress params
func (o *UpdateHostLogsProgressParams) WithHTTPClient(client *http.Client) *UpdateHostLogsProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update host logs progress params
func (o *UpdateHostLogsProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the update host logs progress params
func (o *UpdateHostLogsProgressParams) WithClusterID(clusterID strfmt.UUID) *UpdateHostLogsProgressParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update host logs progress params
func (o *UpdateHostLogsProgressParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithHostID adds the hostID to the update host logs progress params
func (o *UpdateHostLogsProgressParams) WithHostID(hostID strfmt.UUID) *UpdateHostLogsProgressParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the update host logs progress params
func (o *UpdateHostLogsProgressParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithLogsProgressParams adds the logsProgressParams to the update host logs progress params
func (o *UpdateHostLogsProgressParams) WithLogsProgressParams(logsProgressParams *models.LogsProgressParams) *UpdateHostLogsProgressParams {
	o.SetLogsProgressParams(logsProgressParams)
	return o
}

// SetLogsProgressParams adds the logsProgressParams to the update host logs progress params
func (o *UpdateHostLogsProgressParams) SetLogsProgressParams(logsProgressParams *models.LogsProgressParams) {
	o.LogsProgressParams = logsProgressParams
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateHostLogsProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	if o.LogsProgressParams != nil {
		if err := r.SetBodyParam(o.LogsProgressParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}