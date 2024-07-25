package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/redis/resource-manager/Microsoft.Cache/stable/2024-03-01/examples/RedisCacheGet.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.ResourceInfo = armredis.ResourceInfo{
	// 	Name: to.Ptr("cache1"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armredis.Properties{
	// 		EnableNonSSLPort: to.Ptr(true),
	// 		PublicNetworkAccess: to.Ptr(armredis.PublicNetworkAccessEnabled),
	// 		RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
	// 		},
	// 		RedisVersion: to.Ptr("3.2"),
	// 		ReplicasPerMaster: to.Ptr[int32](2),
	// 		ReplicasPerPrimary: to.Ptr[int32](2),
	// 		UpdateChannel: to.Ptr(armredis.UpdateChannelStable),
	// 		SKU: &armredis.SKU{
	// 			Name: to.Ptr(armredis.SKUNamePremium),
	// 			Capacity: to.Ptr[int32](3),
	// 			Family: to.Ptr(armredis.SKUFamilyP),
	// 		},
	// 		HostName: to.Ptr("cache1.redis.cache.windows.net"),
	// 		Instances: []*armredis.InstanceDetails{
	// 			{
	// 				IsMaster: to.Ptr(true),
	// 				IsPrimary: to.Ptr(true),
	// 				NonSSLPort: to.Ptr[int32](13000),
	// 				SSLPort: to.Ptr[int32](15000),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(false),
	// 				IsPrimary: to.Ptr(false),
	// 				NonSSLPort: to.Ptr[int32](13001),
	// 				SSLPort: to.Ptr[int32](15001),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(false),
	// 				IsPrimary: to.Ptr(false),
	// 				NonSSLPort: to.Ptr[int32](13002),
	// 				SSLPort: to.Ptr[int32](15002),
	// 		}},
	// 		LinkedServers: []*armredis.LinkedServer{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/linkedServers/cache2"),
	// 		}},
	// 		Port: to.Ptr[int32](6379),
	// 		PrivateEndpointConnections: []*armredis.PrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/privateEndpointConnections/cachePec"),
	// 				Properties: &armredis.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armredis.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/cachePe"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("Please approve my connection"),
	// 						ActionsRequired: to.Ptr("None"),
	// 						Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
	// 					},
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armredis.ProvisioningStateSucceeded),
	// 		SSLPort: to.Ptr[int32](6380),
	// 	},
	// }
}
