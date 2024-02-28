package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("rg1", "cache1", nil)
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
		// page.PrivateEndpointConnectionListResult = armredisenterprise.PrivateEndpointConnectionListResult{
		// 	Value: []*armredisenterprise.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("pectest01"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/pectest01"),
		// 			Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armredisenterprise.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pectest01"),
		// 			Type: to.Ptr("Microsoft.Cache/redisEnterprise/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/pectest01"),
		// 			Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/petest01"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Auto-Approved"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armredisenterprise.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
