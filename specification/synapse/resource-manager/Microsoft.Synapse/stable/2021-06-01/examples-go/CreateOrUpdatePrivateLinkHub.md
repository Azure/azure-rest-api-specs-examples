Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsynapse%2Farmsynapse%2Fv0.1.0/sdk/resourcemanager/synapse/armsynapse/README.md) on how to add the SDK to your project and authenticate.

```go
package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdatePrivateLinkHub.json
func ExamplePrivateLinkHubsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsynapse.NewPrivateLinkHubsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<private-link-hub-name>",
		armsynapse.PrivateLinkHub{
			TrackedResource: armsynapse.TrackedResource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"key": to.StringPtr("value"),
				},
			},
			Properties: &armsynapse.PrivateLinkHubProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PrivateLinkHub.ID: %s\n", *res.ID)
}
```
