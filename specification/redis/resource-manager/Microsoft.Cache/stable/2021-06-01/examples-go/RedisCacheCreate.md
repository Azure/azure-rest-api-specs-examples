Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredis%2Farmredis%2Fv0.5.0/sdk/resourcemanager/redis/armredis/README.md) on how to add the SDK to your project and authenticate.

```go
package armredis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheCreate.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armredis.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<name>",
		armredis.CreateParameters{
			Location: to.Ptr("<location>"),
			Properties: &armredis.CreateProperties{
				EnableNonSSLPort:  to.Ptr(true),
				MinimumTLSVersion: to.Ptr(armredis.TLSVersionOne2),
				RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
					MaxmemoryPolicy: to.Ptr("<maxmemory-policy>"),
				},
				RedisVersion:       to.Ptr("<redis-version>"),
				ReplicasPerPrimary: to.Ptr[int32](2),
				ShardCount:         to.Ptr[int32](2),
				SKU: &armredis.SKU{
					Name:     to.Ptr(armredis.SKUNamePremium),
					Capacity: to.Ptr[int32](1),
					Family:   to.Ptr(armredis.SKUFamilyP),
				},
				StaticIP: to.Ptr("<static-ip>"),
				SubnetID: to.Ptr("<subnet-id>"),
			},
			Zones: []*string{
				to.Ptr("1")},
		},
		&armredis.ClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
