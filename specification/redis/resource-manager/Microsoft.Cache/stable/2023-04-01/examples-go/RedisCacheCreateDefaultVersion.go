package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/redis/resource-manager/Microsoft.Cache/stable/2023-04-01/examples/RedisCacheCreateDefaultVersion.json
func ExampleClient_BeginCreate_redisCacheCreateDefaultVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredis.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginCreate(ctx, "rg1", "cache1", armredis.CreateParameters{
		Location: to.Ptr("West US"),
		Properties: &armredis.CreateProperties{
			EnableNonSSLPort:  to.Ptr(true),
			MinimumTLSVersion: to.Ptr(armredis.TLSVersionOne2),
			RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
				MaxmemoryPolicy: to.Ptr("allkeys-lru"),
			},
			ReplicasPerPrimary: to.Ptr[int32](2),
			ShardCount:         to.Ptr[int32](2),
			SKU: &armredis.SKU{
				Name:     to.Ptr(armredis.SKUNamePremium),
				Capacity: to.Ptr[int32](1),
				Family:   to.Ptr(armredis.SKUFamilyP),
			},
			StaticIP: to.Ptr("192.168.0.5"),
			SubnetID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/network1/subnets/subnet1"),
		},
		Zones: []*string{
			to.Ptr("1")},
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
	// 		EnableNonSSLPort: to.Ptr(false),
	// 		MinimumTLSVersion: to.Ptr(armredis.TLSVersionOne2),
	// 		RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
	// 			Maxclients: to.Ptr("1000"),
	// 			MaxmemoryDelta: to.Ptr("50"),
	// 			MaxmemoryReserved: to.Ptr("50"),
	// 		},
	// 		RedisVersion: to.Ptr("6.0.14"),
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
	// 				ShardID: to.Ptr[int32](0),
	// 				SSLPort: to.Ptr[int32](15000),
	// 				Zone: to.Ptr("1"),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(false),
	// 				IsPrimary: to.Ptr(false),
	// 				NonSSLPort: to.Ptr[int32](13001),
	// 				ShardID: to.Ptr[int32](0),
	// 				SSLPort: to.Ptr[int32](15001),
	// 				Zone: to.Ptr("1"),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(false),
	// 				IsPrimary: to.Ptr(false),
	// 				NonSSLPort: to.Ptr[int32](13002),
	// 				ShardID: to.Ptr[int32](0),
	// 				SSLPort: to.Ptr[int32](15002),
	// 				Zone: to.Ptr("1"),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(true),
	// 				IsPrimary: to.Ptr(true),
	// 				NonSSLPort: to.Ptr[int32](13003),
	// 				ShardID: to.Ptr[int32](1),
	// 				SSLPort: to.Ptr[int32](15003),
	// 				Zone: to.Ptr("1"),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(false),
	// 				IsPrimary: to.Ptr(false),
	// 				NonSSLPort: to.Ptr[int32](13004),
	// 				ShardID: to.Ptr[int32](1),
	// 				SSLPort: to.Ptr[int32](15004),
	// 				Zone: to.Ptr("1"),
	// 			},
	// 			{
	// 				IsMaster: to.Ptr(false),
	// 				IsPrimary: to.Ptr(false),
	// 				NonSSLPort: to.Ptr[int32](13005),
	// 				ShardID: to.Ptr[int32](1),
	// 				SSLPort: to.Ptr[int32](15005),
	// 				Zone: to.Ptr("1"),
	// 		}},
	// 		Port: to.Ptr[int32](6379),
	// 		ProvisioningState: to.Ptr(armredis.ProvisioningStateSucceeded),
	// 		SSLPort: to.Ptr[int32](6380),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
