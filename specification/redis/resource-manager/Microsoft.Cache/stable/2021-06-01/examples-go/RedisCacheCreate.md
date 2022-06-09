```go
package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheCreate.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredis.NewClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"cache1",
		armredis.CreateParameters{
			Location: to.Ptr("West US"),
			Properties: &armredis.CreateProperties{
				EnableNonSSLPort:  to.Ptr(true),
				MinimumTLSVersion: to.Ptr(armredis.TLSVersionOne2),
				RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
					MaxmemoryPolicy: to.Ptr("allkeys-lru"),
				},
				RedisVersion:       to.Ptr("4"),
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
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredis%2Farmredis%2Fv1.0.0/sdk/resourcemanager/redis/armredis/README.md) on how to add the SDK to your project and authenticate.
