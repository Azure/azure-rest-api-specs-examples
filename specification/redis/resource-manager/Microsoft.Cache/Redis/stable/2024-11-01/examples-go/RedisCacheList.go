package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheList.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
		// page = armredis.ClientListBySubscriptionResponse{
		// 	ListResult: armredis.ListResult{
		// 		Value: []*armredis.ResourceInfo{
		// 			{
		// 				Name: to.Ptr("cache1"),
		// 				Type: to.Ptr("Microsoft.Cache/Redis"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1"),
		// 				Location: to.Ptr("East US"),
		// 				Properties: &armredis.Properties{
		// 					EnableNonSSLPort: to.Ptr(true),
		// 					HostName: to.Ptr("cache1.redis.cache.windows.net"),
		// 					Instances: []*armredis.InstanceDetails{
		// 						{
		// 							IsMaster: to.Ptr(true),
		// 							IsPrimary: to.Ptr(true),
		// 							NonSSLPort: to.Ptr[int32](13000),
		// 							SSLPort: to.Ptr[int32](15000),
		// 						},
		// 						{
		// 							IsMaster: to.Ptr(false),
		// 							IsPrimary: to.Ptr(false),
		// 							NonSSLPort: to.Ptr[int32](13001),
		// 							SSLPort: to.Ptr[int32](15001),
		// 						},
		// 						{
		// 							IsMaster: to.Ptr(false),
		// 							IsPrimary: to.Ptr(false),
		// 							NonSSLPort: to.Ptr[int32](13002),
		// 							SSLPort: to.Ptr[int32](15002),
		// 						},
		// 					},
		// 					Port: to.Ptr[int32](6379),
		// 					ProvisioningState: to.Ptr(armredis.ProvisioningStateCreating),
		// 					RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
		// 					},
		// 					RedisVersion: to.Ptr("3.2"),
		// 					ReplicasPerMaster: to.Ptr[int32](2),
		// 					ReplicasPerPrimary: to.Ptr[int32](2),
		// 					SKU: &armredis.SKU{
		// 						Name: to.Ptr(armredis.SKUNameStandard),
		// 						Capacity: to.Ptr[int32](6),
		// 						Family: to.Ptr(armredis.SKUFamilyC),
		// 					},
		// 					SSLPort: to.Ptr[int32](6380),
		// 					UpdateChannel: to.Ptr(armredis.UpdateChannelStable),
		// 					ZonalAllocationPolicy: to.Ptr(armredis.ZonalAllocationPolicyAutomatic),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
