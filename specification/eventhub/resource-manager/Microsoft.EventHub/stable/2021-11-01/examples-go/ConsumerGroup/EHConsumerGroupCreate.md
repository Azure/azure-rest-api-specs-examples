Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.3.1/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/ConsumerGroup/EHConsumerGroupCreate.json
func ExampleConsumerGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventhub.NewConsumerGroupsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<event-hub-name>",
		"<consumer-group-name>",
		armeventhub.ConsumerGroup{
			Properties: &armeventhub.ConsumerGroupProperties{
				UserMetadata: to.StringPtr("<user-metadata>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ConsumerGroupsClientCreateOrUpdateResult)
}
```
