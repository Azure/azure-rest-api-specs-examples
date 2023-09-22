package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/20312e2b31df58f0ea7560e87062d62aa92f0a14/specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/RedisCacheLinkedServer_List.json
func ExampleLinkedServerClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkedServerClient().NewListPager("rg1", "cache1", nil)
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
		// page.LinkedServerWithPropertiesList = armredis.LinkedServerWithPropertiesList{
		// 	Value: []*armredis.LinkedServerWithProperties{
		// 		{
		// 			Name: to.Ptr("cache2"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/linkedServers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/linkedServers/cache2"),
		// 			Properties: &armredis.LinkedServerProperties{
		// 				GeoReplicatedPrimaryHostName: to.Ptr("cache2.geo.redis.cache.windows.net"),
		// 				LinkedRedisCacheID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache2"),
		// 				LinkedRedisCacheLocation: to.Ptr("West US"),
		// 				PrimaryHostName: to.Ptr("cache1.redis.cache.windows.net"),
		// 				ServerRole: to.Ptr(armredis.ReplicationRoleSecondary),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("cache3"),
		// 			Type: to.Ptr("Microsoft.Cache/Redis/linkedServers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/linkedServers/cache3"),
		// 			Properties: &armredis.LinkedServerProperties{
		// 				GeoReplicatedPrimaryHostName: to.Ptr("cache3.geo.redis.cache.windows.net"),
		// 				LinkedRedisCacheID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache3"),
		// 				LinkedRedisCacheLocation: to.Ptr("West US"),
		// 				PrimaryHostName: to.Ptr("cache1.redis.cache.windows.net"),
		// 				ServerRole: to.Ptr(armredis.ReplicationRoleSecondary),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
