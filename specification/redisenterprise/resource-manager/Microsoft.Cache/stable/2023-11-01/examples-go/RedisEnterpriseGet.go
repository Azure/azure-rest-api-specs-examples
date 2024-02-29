package armredisenterprise_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/redisenterprise/resource-manager/Microsoft.Cache/stable/2023-11-01/examples/RedisEnterpriseGet.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredisenterprise.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().Get(ctx, "rg1", "cache1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Cluster = armredisenterprise.Cluster{
	// 	Name: to.Ptr("cache1"),
	// 	Type: to.Ptr("Microsoft.Cache/redisEnterprise"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armredisenterprise.ClusterProperties{
	// 		HostName: to.Ptr("cache1.westus.something.azure.com"),
	// 		MinimumTLSVersion: to.Ptr(armredisenterprise.TLSVersionOne2),
	// 		PrivateEndpointConnections: []*armredisenterprise.PrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/privateEndpointConnections/cachePec"),
	// 				Properties: &armredisenterprise.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armredisenterprise.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/cachePe"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armredisenterprise.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Please approve my connection"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armredisenterprise.PrivateEndpointServiceConnectionStatusApproved),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armredisenterprise.ProvisioningStateSucceeded),
	// 		RedisVersion: to.Ptr("6"),
	// 		ResourceState: to.Ptr(armredisenterprise.ResourceStateRunning),
	// 	},
	// 	SKU: &armredisenterprise.SKU{
	// 		Name: to.Ptr(armredisenterprise.SKUNameEnterpriseFlashF300),
	// 		Capacity: to.Ptr[int32](3),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1"),
	// 		to.Ptr("2"),
	// 		to.Ptr("3")},
	// 	}
}
