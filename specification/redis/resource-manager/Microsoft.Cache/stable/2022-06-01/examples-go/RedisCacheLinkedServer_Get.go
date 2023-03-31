package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/065033d1c4087a2b009e71c0b3f0666718354ebd/specification/redis/resource-manager/Microsoft.Cache/stable/2022-06-01/examples/RedisCacheLinkedServer_Get.json
func ExampleLinkedServerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkedServerClient().Get(ctx, "rg1", "cache1", "cache2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LinkedServerWithProperties = armredis.LinkedServerWithProperties{
	// 	Name: to.Ptr("cache2"),
	// 	Type: to.Ptr("Microsoft.Cache/Redis/linkedServers"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache1/linkedServers/cache2"),
	// 	Properties: &armredis.LinkedServerProperties{
	// 		GeoReplicatedPrimaryHostName: to.Ptr("cache2.geo.redis.cache.windows.net"),
	// 		LinkedRedisCacheID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache2"),
	// 		LinkedRedisCacheLocation: to.Ptr("West US"),
	// 		PrimaryHostName: to.Ptr("cache1.redis.cache.windows.net"),
	// 		ServerRole: to.Ptr(armredis.ReplicationRoleSecondary),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
