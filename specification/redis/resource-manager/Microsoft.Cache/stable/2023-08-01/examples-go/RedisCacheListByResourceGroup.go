package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/20312e2b31df58f0ea7560e87062d62aa92f0a14/specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheListByResourceGroup.json
func ExampleClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
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
		// 				SKU: &armredis.SKU{
		// 					Name: to.Ptr(armredis.SKUNameStandard),
		// 					Capacity: to.Ptr[int32](6),
		// 					Family: to.Ptr(armredis.SKUFamilyC),
		// 				},
		// 				HostName: to.Ptr("cache1.redis.cache.windows.net"),
		// 				Port: to.Ptr[int32](6379),
		// 				ProvisioningState: to.Ptr(armredis.ProvisioningStateSucceeded),
		// 				SSLPort: to.Ptr[int32](6380),
		// 			},
		// 	}},
		// }
	}
}
