Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredis%2Farmredis%2Fv0.3.1/sdk/resourcemanager/redis/armredis/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheCreate.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredis.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<name>",
		armredis.CreateParameters{
			Location: to.StringPtr("<location>"),
			Properties: &armredis.CreateProperties{
				EnableNonSSLPort:  to.BoolPtr(true),
				MinimumTLSVersion: armredis.TLSVersion("1.2").ToPtr(),
				RedisConfiguration: &armredis.CommonPropertiesRedisConfiguration{
					MaxmemoryPolicy: to.StringPtr("<maxmemory-policy>"),
				},
				RedisVersion:       to.StringPtr("<redis-version>"),
				ReplicasPerPrimary: to.Int32Ptr(2),
				ShardCount:         to.Int32Ptr(2),
				SKU: &armredis.SKU{
					Name:     armredis.SKUName("Premium").ToPtr(),
					Capacity: to.Int32Ptr(1),
					Family:   armredis.SKUFamily("P").ToPtr(),
				},
				StaticIP: to.StringPtr("<static-ip>"),
				SubnetID: to.StringPtr("<subnet-id>"),
			},
			Zones: []*string{
				to.StringPtr("1")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientCreateResult)
}
```
