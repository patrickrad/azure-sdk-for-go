//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnewrelicobservability

import "time"

// AccountInfo - Account Info of the NewRelic account
type AccountInfo struct {
	// Account id
	AccountID *string `json:"accountId,omitempty"`

	// ingestion key of account
	IngestionKey *string `json:"ingestionKey,omitempty"`

	// NewRelic account region
	Region *string `json:"region,omitempty"`
}

// AccountProperties - List of all the New relic accounts for the given user
type AccountProperties struct {
	// account id
	AccountID *string `json:"accountId,omitempty"`

	// account name
	AccountName *string `json:"accountName,omitempty"`

	// organization id
	OrganizationID *string `json:"organizationId,omitempty"`

	// region
	Region *string `json:"region,omitempty"`
}

// AccountPropertiesForNewRelic - Properties of the NewRelic account
type AccountPropertiesForNewRelic struct {
	// NewRelic Account Information
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`

	// NewRelic Organization Information
	OrganizationInfo *OrganizationInfo `json:"organizationInfo,omitempty"`

	// date when plan was applied
	SingleSignOnProperties *NewRelicSingleSignOnProperties `json:"singleSignOnProperties,omitempty"`

	// User id
	UserID *string `json:"userId,omitempty"`
}

// AccountResource - The details of a account resource.
type AccountResource struct {
	// The resource-specific properties for this resource.
	Properties *AccountProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AccountsClientListOptions contains the optional parameters for the AccountsClient.NewListPager method.
type AccountsClientListOptions struct {
	// placeholder for future optional parameters
}

// AccountsListResponse - Response of get all accounts Operation.
type AccountsListResponse struct {
	// REQUIRED; The AccountResource items on this page
	Value []*AccountResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// AppServiceInfo - Details of VM Resource having NewRelic OneAgent installed
type AppServiceInfo struct {
	// Status of the NewRelic agent installed on the App service.
	AgentStatus *string `json:"agentStatus,omitempty"`

	// Version of the NewRelic agent installed on the App service.
	AgentVersion *string `json:"agentVersion,omitempty"`

	// Azure App service resource ID
	AzureResourceID *string `json:"azureResourceId,omitempty"`
}

// AppServicesGetRequest - Request of a app services get Operation.
type AppServicesGetRequest struct {
	// REQUIRED; User Email
	UserEmail *string `json:"userEmail,omitempty"`

	// Azure resource IDs
	AzureResourceIDs []*string `json:"azureResourceIds,omitempty"`
}

// AppServicesListResponse - Response of a list app services Operation.
type AppServicesListResponse struct {
	// REQUIRED; The AppServiceInfo items on this page
	Value []*AppServiceInfo `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them
// from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag. Exclusion takes priority over inclusion.
	Action *TagAction `json:"action,omitempty"`

	// The name (also known as the key) of the tag.
	Name *string `json:"name,omitempty"`

	// The value of the tag.
	Value *string `json:"value,omitempty"`
}

// HostsGetRequest - Request of a Hosts get Operation.
type HostsGetRequest struct {
	// REQUIRED; User Email
	UserEmail *string `json:"userEmail,omitempty"`

	// VM resource IDs
	VMIDs []*string `json:"vmIds,omitempty"`
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty,
	// all resources will be captured. If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
	// resources with the associated tags.
	FilteringTags []*FilteringTag `json:"filteringTags,omitempty"`

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *SendAADLogsStatus `json:"sendAadLogs,omitempty"`

	// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
	SendActivityLogs *SendActivityLogsStatus `json:"sendActivityLogs,omitempty"`

	// Flag specifying if subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *SendSubscriptionLogsStatus `json:"sendSubscriptionLogs,omitempty"`
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType `json:"type,omitempty"`

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// MetricRules - Set of rules for sending metrics for the Monitor resource.
type MetricRules struct {
	// List of filtering tags to be used for capturing metrics.
	FilteringTags []*FilteringTag `json:"filteringTags,omitempty"`

	// Flag specifying if metrics should be sent for the Monitor resource.
	SendMetrics *SendMetricsStatus `json:"sendMetrics,omitempty"`

	// User Email
	UserEmail *string `json:"userEmail,omitempty"`
}

// MetricsRequest - Request of get metrics Operation.
type MetricsRequest struct {
	// REQUIRED; User Email
	UserEmail *string `json:"userEmail,omitempty"`
}

// MetricsStatusRequest - Request of get metrics status Operation.
type MetricsStatusRequest struct {
	// REQUIRED; User Email
	UserEmail *string `json:"userEmail,omitempty"`

	// Azure resource IDs
	AzureResourceIDs []*string `json:"azureResourceIds,omitempty"`
}

// MetricsStatusResponse - Response of get metrics status Operation.
type MetricsStatusResponse struct {
	// Azure resource IDs
	AzureResourceIDs []*string `json:"azureResourceIds,omitempty"`
}

// MonitorProperties - Properties specific to the NewRelic Monitor resource
type MonitorProperties struct {
	// Source of account creation
	AccountCreationSource *AccountCreationSource `json:"accountCreationSource,omitempty"`

	// MarketplaceSubscriptionStatus of the resource
	NewRelicAccountProperties *AccountPropertiesForNewRelic `json:"newRelicAccountProperties,omitempty"`

	// Source of org creation
	OrgCreationSource *OrgCreationSource `json:"orgCreationSource,omitempty"`

	// Plan details
	PlanData *PlanData `json:"planData,omitempty"`

	// User Info
	UserInfo *UserInfo `json:"userInfo,omitempty"`

	// READ-ONLY; Liftr resource category
	LiftrResourceCategory *LiftrResourceCategories `json:"liftrResourceCategory,omitempty" azure:"ro"`

	// READ-ONLY; Liftr resource preference. The priority of the resource.
	LiftrResourcePreference *int32 `json:"liftrResourcePreference,omitempty" azure:"ro"`

	// READ-ONLY; Marketplace Subscription Id
	MarketplaceSubscriptionID *string `json:"marketplaceSubscriptionId,omitempty" azure:"ro"`

	// READ-ONLY; NewRelic Organization properties of the resource
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus `json:"marketplaceSubscriptionStatus,omitempty" azure:"ro"`

	// READ-ONLY; MonitoringStatus of the resource
	MonitoringStatus *MonitoringStatus `json:"monitoringStatus,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning State of the resource
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MonitoredResource - Details of resource being monitored by NewRelic monitor resource
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string `json:"id,omitempty"`

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string `json:"reasonForLogsStatus,omitempty"`

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string `json:"reasonForMetricsStatus,omitempty"`

	// Flag indicating if resource is sending logs to NewRelic.
	SendingLogs *SendingLogsStatus `json:"sendingLogs,omitempty"`

	// Flag indicating if resource is sending metrics to NewRelic.
	SendingMetrics *SendingMetricsStatus `json:"sendingMetrics,omitempty"`
}

// MonitoredResourceListResponse - List of all the resources being monitored by NewRelic monitor resource
type MonitoredResourceListResponse struct {
	// REQUIRED; The MonitoredResource items on this page
	Value []*MonitoredResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// MonitoringTagRulesProperties - The resource-specific properties for this resource.
type MonitoringTagRulesProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules `json:"logRules,omitempty"`

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules `json:"metricRules,omitempty"`

	// READ-ONLY; Provisioning State of the resource
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MonitorsClientBeginCreateOrUpdateOptions contains the optional parameters for the MonitorsClient.BeginCreateOrUpdate method.
type MonitorsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
type MonitorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientGetMetricRulesOptions contains the optional parameters for the MonitorsClient.GetMetricRules method.
type MonitorsClientGetMetricRulesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientGetMetricStatusOptions contains the optional parameters for the MonitorsClient.GetMetricStatus method.
type MonitorsClientGetMetricStatusOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
type MonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListAppServicesOptions contains the optional parameters for the MonitorsClient.NewListAppServicesPager method.
type MonitorsClientListAppServicesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.NewListByResourceGroupPager
// method.
type MonitorsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListBySubscriptionOptions contains the optional parameters for the MonitorsClient.NewListBySubscriptionPager
// method.
type MonitorsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListHostsOptions contains the optional parameters for the MonitorsClient.NewListHostsPager method.
type MonitorsClientListHostsOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListMonitoredResourcesOptions contains the optional parameters for the MonitorsClient.NewListMonitoredResourcesPager
// method.
type MonitorsClientListMonitoredResourcesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientSwitchBillingOptions contains the optional parameters for the MonitorsClient.SwitchBilling method.
type MonitorsClientSwitchBillingOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientUpdateOptions contains the optional parameters for the MonitorsClient.Update method.
type MonitorsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientVMHostPayloadOptions contains the optional parameters for the MonitorsClient.VMHostPayload method.
type MonitorsClientVMHostPayloadOptions struct {
	// placeholder for future optional parameters
}

// NewRelicMonitorResource - A Monitor Resource by NewRelic
type NewRelicMonitorResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitorProperties `json:"properties,omitempty"`

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// NewRelicMonitorResourceListResult - The response of a NewRelicMonitorResource list operation.
type NewRelicMonitorResourceListResult struct {
	// REQUIRED; The NewRelicMonitorResource items on this page
	Value []*NewRelicMonitorResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// NewRelicMonitorResourceUpdate - The type used for update operations of the NewRelicMonitorResource.
type NewRelicMonitorResourceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity `json:"identity,omitempty"`

	// The updatable properties of the NewRelicMonitorResource.
	Properties *NewRelicMonitorResourceUpdateProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// NewRelicMonitorResourceUpdateProperties - The updatable properties of the NewRelicMonitorResource.
type NewRelicMonitorResourceUpdateProperties struct {
	// Source of account creation
	AccountCreationSource *AccountCreationSource `json:"accountCreationSource,omitempty"`

	// MarketplaceSubscriptionStatus of the resource
	NewRelicAccountProperties *AccountPropertiesForNewRelic `json:"newRelicAccountProperties,omitempty"`

	// Source of org creation
	OrgCreationSource *OrgCreationSource `json:"orgCreationSource,omitempty"`

	// Plan details
	PlanData *PlanData `json:"planData,omitempty"`

	// User Info
	UserInfo *UserInfo `json:"userInfo,omitempty"`
}

// NewRelicSingleSignOnProperties - Single sign on Info of the NewRelic account
type NewRelicSingleSignOnProperties struct {
	// The Id of the Enterprise App used for Single sign-on.
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty"`

	// Provisioning state
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`

	// Single sign-on state
	SingleSignOnState *SingleSignOnStates `json:"singleSignOnState,omitempty"`

	// The login URL specific to this NewRelic Organization
	SingleSignOnURL *string `json:"singleSignOnUrl,omitempty"`
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OrganizationInfo - Organization Info of the NewRelic account
type OrganizationInfo struct {
	// Organization id
	OrganizationID *string `json:"organizationId,omitempty"`
}

// OrganizationProperties - Details of Organizations
type OrganizationProperties struct {
	// Billing source
	BillingSource *BillingSource `json:"billingSource,omitempty"`

	// organization id
	OrganizationID *string `json:"organizationId,omitempty"`

	// organization name
	OrganizationName *string `json:"organizationName,omitempty"`
}

// OrganizationResource - The details of a Organization resource.
type OrganizationResource struct {
	// The resource-specific properties for this resource.
	Properties *OrganizationProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// OrganizationsClientListOptions contains the optional parameters for the OrganizationsClient.NewListPager method.
type OrganizationsClientListOptions struct {
	// placeholder for future optional parameters
}

// OrganizationsListResponse - Response of get all organizations Operation.
type OrganizationsListResponse struct {
	// REQUIRED; The OrganizationResource items on this page
	Value []*OrganizationResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// PlanData - Plan data of NewRelic Monitor resource
type PlanData struct {
	// Different billing cycles like MONTHLY/WEEKLY. this could be enum
	BillingCycle *BillingCycle `json:"billingCycle,omitempty"`

	// date when plan was applied
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// plan id as published by NewRelic
	PlanDetails *string `json:"planDetails,omitempty"`

	// Different usage type like PAYG/COMMITTED. this could be enum
	UsageType *UsageType `json:"usageType,omitempty"`
}

// PlanDataListResponse - Response of get all plan data Operation.
type PlanDataListResponse struct {
	// REQUIRED; The PlanDataResource items on this page
	Value []*PlanDataResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// PlanDataProperties - Plan details
type PlanDataProperties struct {
	// Source of account creation
	AccountCreationSource *AccountCreationSource `json:"accountCreationSource,omitempty"`

	// Source of org creation
	OrgCreationSource *OrgCreationSource `json:"orgCreationSource,omitempty"`

	// Plan details
	PlanData *PlanData `json:"planData,omitempty"`
}

// PlanDataResource - The details of a PlanData resource.
type PlanDataResource struct {
	// The resource-specific properties for this resource.
	Properties *PlanDataProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// PlansClientListOptions contains the optional parameters for the PlansClient.NewListPager method.
type PlansClientListOptions struct {
	// Account Id.
	AccountID *string
	// Organization Id.
	OrganizationID *string
}

// SwitchBillingRequest - Request of a switch billing Operation.
type SwitchBillingRequest struct {
	// REQUIRED; User Email
	UserEmail *string `json:"userEmail,omitempty"`

	// Azure resource Id
	AzureResourceID *string `json:"azureResourceId,omitempty"`

	// Organization id
	OrganizationID *string `json:"organizationId,omitempty"`

	// Plan details
	PlanData *PlanData `json:"planData,omitempty"`
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

// TagRule - A tag rule belonging to NewRelic account
type TagRule struct {
	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitoringTagRulesProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TagRuleListResult - The response of a TagRule list operation.
type TagRuleListResult struct {
	// REQUIRED; The TagRule items on this page
	Value []*TagRule `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// TagRuleUpdate - The type used for update operations of the TagRule.
type TagRuleUpdate struct {
	// The updatable properties of the TagRule.
	Properties *TagRuleUpdateProperties `json:"properties,omitempty"`
}

// TagRuleUpdateProperties - The updatable properties of the TagRule.
type TagRuleUpdateProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules `json:"logRules,omitempty"`

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules `json:"metricRules,omitempty"`
}

// TagRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.BeginCreateOrUpdate method.
type TagRulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TagRulesClientBeginDeleteOptions contains the optional parameters for the TagRulesClient.BeginDelete method.
type TagRulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
type TagRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientListByNewRelicMonitorResourceOptions contains the optional parameters for the TagRulesClient.NewListByNewRelicMonitorResourcePager
// method.
type TagRulesClientListByNewRelicMonitorResourceOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientUpdateOptions contains the optional parameters for the TagRulesClient.Update method.
type TagRulesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string `json:"clientId,omitempty" azure:"ro"`

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`
}

// UserInfo - User Info of NewRelic Monitor resource
type UserInfo struct {
	// country if user
	Country *string `json:"country,omitempty"`

	// User Email
	EmailAddress *string `json:"emailAddress,omitempty"`

	// First name
	FirstName *string `json:"firstName,omitempty"`

	// Last name
	LastName *string `json:"lastName,omitempty"`

	// Contact phone number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// VMExtensionPayload - Response of payload to be passed while installing VM agent.
type VMExtensionPayload struct {
	// Ingestion key of the account
	IngestionKey *string `json:"ingestionKey,omitempty"`
}

// VMHostsListResponse - Response of a list VM Host Operation.
type VMHostsListResponse struct {
	// REQUIRED; The VMInfo items on this page
	Value []*VMInfo `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// VMInfo - Details of VM Resource having NewRelic OneAgent installed
type VMInfo struct {
	// Status of the NewRelic agent installed on the VM.
	AgentStatus *string `json:"agentStatus,omitempty"`

	// Version of the NewRelic agent installed on the VM.
	AgentVersion *string `json:"agentVersion,omitempty"`

	// Azure VM resource ID
	VMID *string `json:"vmId,omitempty"`
}
