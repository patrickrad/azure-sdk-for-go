//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconfidentialledger

import "time"

// AADBasedSecurityPrincipal - AAD based security principal with associated Ledger RoleName
type AADBasedSecurityPrincipal struct {
	// LedgerRole associated with the Security Principal of Ledger
	LedgerRoleName *LedgerRoleName `json:"ledgerRoleName,omitempty"`

	// UUID/GUID based Principal Id of the Security Principal
	PrincipalID *string `json:"principalId,omitempty"`

	// UUID/GUID based Tenant Id of the Security Principal
	TenantID *string `json:"tenantId,omitempty"`
}

// CertBasedSecurityPrincipal - Cert based security principal with Ledger RoleName
type CertBasedSecurityPrincipal struct {
	// Public key of the user cert (.pem or .cer)
	Cert *string `json:"cert,omitempty"`

	// LedgerRole associated with the Security Principal of Ledger
	LedgerRoleName *LedgerRoleName `json:"ledgerRoleName,omitempty"`
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string `json:"name,omitempty"`

	// The resource type.
	Type *string `json:"type,omitempty"`
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string `json:"message,omitempty"`

	// Indicates if the resource name is available.
	NameAvailable *bool `json:"nameAvailable,omitempty"`

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason `json:"reason,omitempty"`
}

// ClientCheckNameAvailabilityOptions contains the optional parameters for the Client.CheckNameAvailability method.
type ClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ConfidentialLedger - Confidential Ledger. Contains the properties of Confidential Ledger Resource.
type ConfidentialLedger struct {
	// The Azure location where the Confidential Ledger is running.
	Location *string `json:"location,omitempty"`

	// Properties of Confidential Ledger Resource.
	Properties *LedgerProperties `json:"properties,omitempty"`

	// Additional tags for Confidential Ledger
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the Resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorDetail `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail `json:"error,omitempty"`
}

// LedgerClientBeginCreateOptions contains the optional parameters for the LedgerClient.BeginCreate method.
type LedgerClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LedgerClientBeginDeleteOptions contains the optional parameters for the LedgerClient.BeginDelete method.
type LedgerClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LedgerClientBeginUpdateOptions contains the optional parameters for the LedgerClient.BeginUpdate method.
type LedgerClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// LedgerClientGetOptions contains the optional parameters for the LedgerClient.Get method.
type LedgerClientGetOptions struct {
	// placeholder for future optional parameters
}

// LedgerClientListByResourceGroupOptions contains the optional parameters for the LedgerClient.NewListByResourceGroupPager
// method.
type LedgerClientListByResourceGroupOptions struct {
	// The filter to apply on the list operation. eg. $filter=ledgerType eq 'Public'
	Filter *string
}

// LedgerClientListBySubscriptionOptions contains the optional parameters for the LedgerClient.NewListBySubscriptionPager
// method.
type LedgerClientListBySubscriptionOptions struct {
	// The filter to apply on the list operation. eg. $filter=ledgerType eq 'Public'
	Filter *string
}

// LedgerProperties - Additional Confidential Ledger properties.
type LedgerProperties struct {
	// Array of all AAD based Security Principals.
	AADBasedSecurityPrincipals []*AADBasedSecurityPrincipal `json:"aadBasedSecurityPrincipals,omitempty"`

	// Array of all cert based Security Principals.
	CertBasedSecurityPrincipals []*CertBasedSecurityPrincipal `json:"certBasedSecurityPrincipals,omitempty"`

	// Type of Confidential Ledger
	LedgerType *LedgerType `json:"ledgerType,omitempty"`

	// READ-ONLY; Endpoint for accessing network identity.
	IdentityServiceURI *string `json:"identityServiceUri,omitempty" azure:"ro"`

	// READ-ONLY; Internal namespace for the Ledger
	LedgerInternalNamespace *string `json:"ledgerInternalNamespace,omitempty" azure:"ro"`

	// READ-ONLY; Unique name for the Confidential Ledger.
	LedgerName *string `json:"ledgerName,omitempty" azure:"ro"`

	// READ-ONLY; Endpoint for calling Ledger Service.
	LedgerURI *string `json:"ledgerUri,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of Ledger Resource
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// List - Object that includes an array of Confidential Ledgers and a possible link for next set.
type List struct {
	// The URL the client should use to fetch the next page (per server side paging).
	NextLink *string `json:"nextLink,omitempty"`

	// List of Confidential Ledgers
	Value []*ConfidentialLedger `json:"value,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Resource - An Azure resource.
type Resource struct {
	// READ-ONLY; Fully qualified resource Id for the resource.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of the Resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceLocation - Location of the ARM Resource
type ResourceLocation struct {
	// The Azure location where the Confidential Ledger is running.
	Location *string `json:"location,omitempty"`
}

// ResourceProviderOperationDefinition - Describes the Resource Provider Operation.
type ResourceProviderOperationDefinition struct {
	// Details about the operations
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation is data action or not.
	IsDataAction *bool `json:"isDataAction,omitempty"`

	// Resource provider operation name.
	Name *string `json:"name,omitempty"`
}

// ResourceProviderOperationDisplay - Describes the properties of the Operation.
type ResourceProviderOperationDisplay struct {
	// Description of the resource provider operation.
	Description *string `json:"description,omitempty"`

	// Name of the resource provider operation.
	Operation *string `json:"operation,omitempty"`

	// Name of the resource provider.
	Provider *string `json:"provider,omitempty"`

	// Name of the resource type.
	Resource *string `json:"resource,omitempty"`
}

// ResourceProviderOperationList - List containing this Resource Provider's available operations.
type ResourceProviderOperationList struct {
	// READ-ONLY; The URI that can be used to request the next page for list of Azure operations.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; Resource provider operations list.
	Value []*ResourceProviderOperationDefinition `json:"value,omitempty" azure:"ro"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// Tags for Confidential Ledger Resource
type Tags struct {
	// Additional tags for Confidential Ledger
	Tags map[string]*string `json:"tags,omitempty"`
}
