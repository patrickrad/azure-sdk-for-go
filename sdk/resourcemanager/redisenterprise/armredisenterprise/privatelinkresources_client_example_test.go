//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9565a97e3efbeb3691c9100d5473b0a884c1b786/specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2023-03-01-preview/examples/RedisEnterpriseListPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateLinkResourcesClient().NewListByClusterPager("rg1", "cache1", nil)
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
		// page.PrivateLinkResourceListResult = armredisenterprise.PrivateLinkResourceListResult{
		// 	Value: []*armredisenterprise.PrivateLinkResource{
		// 		{
		// 			Name: to.Ptr("redisEnterpriseCache"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise/privateLinkResources"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateLinkResources/redisEnterpriseCache"),
		// 			Properties: &armredisenterprise.PrivateLinkResourceProperties{
		// 				GroupID: to.Ptr("redisEnterpriseCache"),
		// 				RequiredMembers: []*string{
		// 					to.Ptr("redisEnterpriseCache")},
		// 					RequiredZoneNames: []*string{
		// 						to.Ptr("privatelink.redisenterprise.cache.windows.net")},
		// 					},
		// 			}},
		// 		}
	}
}
