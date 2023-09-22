package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/20312e2b31df58f0ea7560e87062d62aa92f0a14/specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheList.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListBySubscriptionPager(nil)
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
		// page.ListResult = armredis.ListResult{
		// 	Value: []*armredis.ResourceInfo{
		// 		{
		// 			Name: to.Ptr("cache1"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armredis.Properties{
		// 				EnableNonSSLPort: to.Ptr(true),
		// 				RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
		// 				},
		// 				RedisVersion: to.Ptr("3.2"),
		// 				ReplicasPerMaster: to.Ptr[int32](2),
		// 				ReplicasPerPrimary: to.Ptr[int32](2),
		// 				UpdateChannel: to.Ptr(armredis.UpdateChannelStable),
		// 				SKU: &armredis.SKU{
		// 					Name: to.Ptr(armredis.SKUNameStandard),
		// 					Capacity: to.Ptr[int32](6),
		// 					Family: to.Ptr(armredis.SKUFamilyC),
		// 				},
		// 				HostName: to.Ptr("cache1.redis.cache.windows.net"),
		// 				Instances: []*armredis.InstanceDetails{
		// 					{
		// 						IsMaster: to.Ptr(true),
		// 						IsPrimary: to.Ptr(true),
		// 						NonSSLPort: to.Ptr[int32](13000),
		// 						SSLPort: to.Ptr[int32](15000),
		// 					},
		// 					{
		// 						IsMaster: to.Ptr(false),
		// 						IsPrimary: to.Ptr(false),
		// 						NonSSLPort: to.Ptr[int32](13001),
		// 						SSLPort: to.Ptr[int32](15001),
		// 					},
		// 					{
		// 						IsMaster: to.Ptr(false),
		// 						IsPrimary: to.Ptr(false),
		// 						NonSSLPort: to.Ptr[int32](13002),
		// 						SSLPort: to.Ptr[int32](15002),
		// 				}},
		// 				Port: to.Ptr[int32](6379),
		// 				ProvisioningState: to.Ptr(armredis.ProvisioningStateSucceeded),
		// 				SSLPort: to.Ptr[int32](6380),
		// 			},
		// 	}},
		// }
	}
}
