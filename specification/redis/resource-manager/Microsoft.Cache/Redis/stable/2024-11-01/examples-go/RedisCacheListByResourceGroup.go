package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v4"
)

// Generated from example definition: 2024-11-01/RedisCacheListByResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListByResourceGroupPager("rg1", nil)
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
		// page = armredis.ClientListByResourceGroupResponse{
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
		// 					Port: to.Ptr[int32](6379),
		// 					ProvisioningState: to.Ptr(armredis.ProvisioningStateCreating),
		// 					RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
		// 					},
		// 					RedisVersion: to.Ptr("3.2"),
		// 					SKU: &armredis.SKU{
		// 						Name: to.Ptr(armredis.SKUNameStandard),
		// 						Capacity: to.Ptr[int32](6),
		// 						Family: to.Ptr(armredis.SKUFamilyC),
		// 					},
		// 					SSLPort: to.Ptr[int32](6380),
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
