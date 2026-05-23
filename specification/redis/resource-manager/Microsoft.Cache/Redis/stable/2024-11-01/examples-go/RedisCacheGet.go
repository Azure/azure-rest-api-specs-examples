package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheGet.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
	// res = armredis.ClientGetResponse{
	// 	ResourceInfo: &armredis.ResourceInfo{
	// 		Name: to.Ptr("cache1"),
	// 		Type: to.Ptr("Microsoft.Cache/Redis"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armredis.Properties{
	// 			EnableNonSSLPort: to.Ptr(true),
	// 			HostName: to.Ptr("cache1.redis.cache.windows.net"),
	// 			Instances: []*armredis.InstanceDetails{
	// 				{
	// 					IsMaster: to.Ptr(true),
	// 					IsPrimary: to.Ptr(true),
	// 					NonSSLPort: to.Ptr[int32](13000),
	// 					SSLPort: to.Ptr[int32](15000),
	// 				},
	// 				{
	// 					IsMaster: to.Ptr(false),
	// 					IsPrimary: to.Ptr(false),
	// 					NonSSLPort: to.Ptr[int32](13001),
	// 					SSLPort: to.Ptr[int32](15001),
	// 				},
	// 				{
	// 					IsMaster: to.Ptr(false),
	// 					IsPrimary: to.Ptr(false),
	// 					NonSSLPort: to.Ptr[int32](13002),
	// 					SSLPort: to.Ptr[int32](15002),
	// 				},
	// 			},
	// 			LinkedServers: []*armredis.LinkedServer{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/linkedServers/cache2"),
	// 				},
	// 			},
	// 			Port: to.Ptr[int32](6379),
	// 			PrivateEndpointConnections: []*armredis.PrivateEndpointConnection{
	// 				{
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/privateEndpointConnections/cachePec"),
	// 					Properties: &armredis.PrivateEndpointConnectionProperties{
	// 						PrivateEndpoint: &armredis.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/cachePe"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armredis.PrivateLinkServiceConnectionState{
	// 							Description: to.Ptr("Please approve my connection"),
	// 							ActionsRequired: to.Ptr("None"),
	// 							Status: to.Ptr(armredis.PrivateEndpointServiceConnectionStatusApproved),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armredis.ProvisioningStateCreating),
	// 			PublicNetworkAccess: to.Ptr(armredis.PublicNetworkAccessEnabled),
	// 			RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
	// 			},
	// 			RedisVersion: to.Ptr("3.2"),
	// 			ReplicasPerMaster: to.Ptr[int32](2),
	// 			ReplicasPerPrimary: to.Ptr[int32](2),
	// 			SKU: &armredis.SKU{
	// 				Name: to.Ptr(armredis.SKUNamePremium),
	// 				Capacity: to.Ptr[int32](3),
	// 				Family: to.Ptr(armredis.SKUFamilyP),
	// 			},
	// 			SSLPort: to.Ptr[int32](6380),
	// 			UpdateChannel: to.Ptr(armredis.UpdateChannelStable),
	// 			ZonalAllocationPolicy: to.Ptr(armredis.ZonalAllocationPolicyAutomatic),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
