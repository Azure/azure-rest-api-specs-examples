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

// x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2021-06-01/examples/RedisCacheLinkedServer_Create.json
func ExampleLinkedServerClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armredis.NewLinkedServerClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<name>",
		"<linked-server-name>",
		armredis.LinkedServerCreateParameters{
			Properties: &armredis.LinkedServerCreateProperties{
				LinkedRedisCacheID:       to.StringPtr("<linked-redis-cache-id>"),
				LinkedRedisCacheLocation: to.StringPtr("<linked-redis-cache-location>"),
				ServerRole:               armredis.ReplicationRoleSecondary.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LinkedServerClientCreateResult)
}
```
