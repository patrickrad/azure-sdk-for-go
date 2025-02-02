//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf3574813e15bb33b3cb610f44edfcbebd8b1b23/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/preview/2021-12-01-preview/examples/OperationsListByTenant.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.OperationListResult = armoperationalinsights.OperationListResult{
		// 	Value: []*armoperationalinsights.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/write"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new workspace or links to an existing workspace by providing the customer id from the existing workspace."),
		// 				Operation: to.Ptr("Create Workspace"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Workspace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets an existing workspace"),
		// 				Operation: to.Ptr("Get Workspace"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Workspace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/delete"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes a workspace. If the workspace was linked to an existing workspace at creation time then the workspace it was linked to is not deleted."),
		// 				Operation: to.Ptr("Delete Workspace"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Workspace"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/generateregistrationcertificate/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Generates Registration Certificate for the workspace. This Certificate is used to connect Microsoft System Center Operation Manager to the workspace."),
		// 				Operation: to.Ptr("Generates Registration Certificate for Workspace."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Registration Certificate"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/storageinsightconfigs/write"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Creates a new storage configuration. These configurations are used to pull data from a location in an existing storage account."),
		// 				Operation: to.Ptr("Create Storage Configuration"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Storage Insight Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/storageinsightconfigs/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets a storage configuration."),
		// 				Operation: to.Ptr("Get Storage Configuration"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Storage Insight Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/storageinsightconfigs/delete"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Deletes a storage configuration. This will stop Microsoft Operational Insights from reading data from the storage account."),
		// 				Operation: to.Ptr("Delete Storage Configuration"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Storage Insight Configuration"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/register/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Register a subscription to a resource provider."),
		// 				Operation: to.Ptr("Register a subscription to a resource provider."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Register"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/sharedKeys/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Retrieves the shared keys for the workspace. These keys are used to connect Microsoft Operational Insights agents to the workspace."),
		// 				Operation: to.Ptr("List Workspace Shared Keys"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Shared Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/sharedKeys/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Retrieves the shared keys for the workspace. These keys are used to connect Microsoft Operational Insights agents to the workspace."),
		// 				Operation: to.Ptr("List Workspace Shared Keys"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Shared Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/listKeys/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Retrieves the list keys for the workspace. These keys are used to connect Microsoft Operational Insights agents to the workspace."),
		// 				Operation: to.Ptr("List Workspace Keys"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("List Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/listKeys/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Retrieves the list keys for the workspace. These keys are used to connect Microsoft Operational Insights agents to the workspace."),
		// 				Operation: to.Ptr("List Workspace Keys"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("List Keys"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/managementGroups/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the names and metadata for System Center Operations Manager management groups connected to this workspace."),
		// 				Operation: to.Ptr("Get Management Groups for Workspace"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Management Group"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/usages/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets usage data for a workspace including the amount of data read by the workspace."),
		// 				Operation: to.Ptr("Get Usage Data for Workspace"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Usage Metric"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/search/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Executes a search query"),
		// 				Operation: to.Ptr("Search Workspace Data"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Search"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/schema/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Gets the search schema for the workspace.  Search schema includes the exposed fields and their types."),
		// 				Operation: to.Ptr("Get Search Schema"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Search Schema"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/datasources/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Get datasources under a workspace."),
		// 				Operation: to.Ptr("Get datasources under a workspace."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Data Source"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/datasources/write"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Create/Update datasources under a workspace."),
		// 				Operation: to.Ptr("Create/Update datasources under a workspace."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Data Source"),
		// 			},
		// 		},
		// 		{
		// 		},
		// 		{
		// 		},
		// 		{
		// 		},
		// 		{
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/notificationSettings/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Get the user's notification settings for the workspace."),
		// 				Operation: to.Ptr("Get Notification Settings"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Notification Settings"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/notificationSettings/write"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Set the user's notification settings for the workspace."),
		// 				Operation: to.Ptr("Put Notification Settings"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Notification Settings"),
		// 			},
		// 		},
		// 		{
		// 		},
		// 		{
		// 		},
		// 		{
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/configurationScopes/delete"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Delete Configuration Scope"),
		// 				Operation: to.Ptr("Delete Configuration Scope"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Configuration Scope"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Get linked services under given workspace."),
		// 				Operation: to.Ptr("Get linked services under given workspace."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Linked Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices/write"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Create/Update linked services under given workspace."),
		// 				Operation: to.Ptr("Create/Update linked services under given workspace."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Linked Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedServices/delete"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Delete linked services under given workspace."),
		// 				Operation: to.Ptr("Delete linked services under given workspace."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Linked Services"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/intelligencepacks/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Lists all intelligence packs that are visible for a given worksapce and also lists whether the pack is enabled or disabled for that workspace."),
		// 				Operation: to.Ptr("List Intelligence Packs"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Intelligence Packs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/intelligencepacks/enable/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Enables an intelligence pack for a given workspace."),
		// 				Operation: to.Ptr("Enable Intelligence Pack"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Intelligence Packs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/intelligencepacks/disable/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Disables an intelligence pack for a given workspace."),
		// 				Operation: to.Ptr("Disable Intelligence Pack"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Intelligence Packs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/analytics/query/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Search using new engine."),
		// 				Operation: to.Ptr("Search using new engine."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("analytics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/analytics/query/schema/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Get search schema V2."),
		// 				Operation: to.Ptr("Get search schema V2."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("analytics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/api/query/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Search using new engine."),
		// 				Operation: to.Ptr("Search using new engine."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("analytics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/api/query/schema/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Get search schema V2."),
		// 				Operation: to.Ptr("Get search schema V2."),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("analytics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/purge/action"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Delete specified data from workspace"),
		// 				Operation: to.Ptr("Delete specified data from workspace"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("analytics"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/linkTargets/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Lists existing accounts that are not associated with an Azure subscription. To link this Azure subscription to a workspace, use a customer id returned by this operation in the customer id property of the Create Workspace operation."),
		// 				Operation: to.Ptr("List Unlinked Accounts"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Unlinked Account"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.OperationalInsights/workspaces/metricDefinitions/read"),
		// 			Display: &armoperationalinsights.OperationDisplay{
		// 				Description: to.Ptr("Get Metric Definitions under workspace"),
		// 				Operation: to.Ptr("Metric Definition operation"),
		// 				Provider: to.Ptr("Microsoft Operational Insights"),
		// 				Resource: to.Ptr("Metric Definitions"),
		// 			},
		// 	}},
		// }
	}
}
