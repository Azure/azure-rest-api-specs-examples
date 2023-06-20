package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheUpdate.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginUpdate(ctx, "rg1", "cache1", armredis.UpdateParameters{
		Properties: &armredis.UpdateProperties{
			EnableNonSSLPort:   to.Ptr(true),
			ReplicasPerPrimary: to.Ptr[int32](2),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
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
	// 		RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
	// 			Maxclients: to.Ptr("1000"),
	// 			MaxmemoryDelta: to.Ptr("50"),
	// 			MaxmemoryReserved: to.Ptr("50"),
	// 		},
	// 		RedisVersion: to.Ptr("3.0"),
	// 		ReplicasPerMaster: to.Ptr[int32](2),
	// 		ReplicasPerPrimary: to.Ptr[int32](2),
	// 		SKU: &armredis.SKU{
	// 			Name: to.Ptr(armredis.SKUNamePremium),
	// 			Capacity: to.Ptr[int32](1),
	// 			Family: to.Ptr(armredis.SKUFamilyP),
	// 		},
	// 		AccessKeys: &armredis.AccessKeys{
	// 			PrimaryKey: to.Ptr("<primaryKey>"),
	// 			SecondaryKey: to.Ptr("<secondaryKey>"),
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
	// 		Port: to.Ptr[int32](6379),
	// 		ProvisioningState: to.Ptr(armredis.ProvisioningStateSucceeded),
	// 		SSLPort: to.Ptr[int32](6380),
	// 	},
	// }
}
