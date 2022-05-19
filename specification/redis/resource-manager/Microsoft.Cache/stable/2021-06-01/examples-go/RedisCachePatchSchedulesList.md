Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fredis%2Farmredis%2Fv1.0.0/sdk/resourcemanager/redis/armredis/README.md) on how to add the SDK to your project and authenticate.

```go
package armredis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCachePatchSchedulesList.json
func ExamplePatchSchedulesClient_NewListByRedisResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armredis.NewPatchSchedulesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByRedisResourcePager("rg1",
		"cache1",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
