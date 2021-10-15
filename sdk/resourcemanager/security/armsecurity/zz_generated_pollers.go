//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// AdaptiveNetworkHardeningsEnforcePoller provides polling facilities until the operation reaches a terminal state.
type AdaptiveNetworkHardeningsEnforcePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *AdaptiveNetworkHardeningsEnforcePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *AdaptiveNetworkHardeningsEnforcePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final AdaptiveNetworkHardeningsEnforceResponse will be returned.
func (p *AdaptiveNetworkHardeningsEnforcePoller) FinalResponse(ctx context.Context) (AdaptiveNetworkHardeningsEnforceResponse, error) {
	respType := AdaptiveNetworkHardeningsEnforceResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return AdaptiveNetworkHardeningsEnforceResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *AdaptiveNetworkHardeningsEnforcePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// AlertsSimulatePoller provides polling facilities until the operation reaches a terminal state.
type AlertsSimulatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *AlertsSimulatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *AlertsSimulatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final AlertsSimulateResponse will be returned.
func (p *AlertsSimulatePoller) FinalResponse(ctx context.Context) (AlertsSimulateResponse, error) {
	respType := AlertsSimulateResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return AlertsSimulateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *AlertsSimulatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// ServerVulnerabilityAssessmentDeletePoller provides polling facilities until the operation reaches a terminal state.
type ServerVulnerabilityAssessmentDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *ServerVulnerabilityAssessmentDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *ServerVulnerabilityAssessmentDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final ServerVulnerabilityAssessmentDeleteResponse will be returned.
func (p *ServerVulnerabilityAssessmentDeletePoller) FinalResponse(ctx context.Context) (ServerVulnerabilityAssessmentDeleteResponse, error) {
	respType := ServerVulnerabilityAssessmentDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return ServerVulnerabilityAssessmentDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *ServerVulnerabilityAssessmentDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
