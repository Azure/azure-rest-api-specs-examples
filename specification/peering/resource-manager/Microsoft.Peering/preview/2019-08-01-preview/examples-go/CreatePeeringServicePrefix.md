Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpeering%2Farmpeering%2Fv0.2.1/sdk/resourcemanager/peering/armpeering/README.md) on how to add the SDK to your project and authenticate.

```go
package armpeering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering"
)

// x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/preview/2019-08-01-preview/examples/CreatePeeringServicePrefix.json
func ExampleServicePrefixesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpeering.NewServicePrefixesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<peering-service-name>",
		"<prefix-name>",
		armpeering.ServicePrefix{
			Properties: &armpeering.ServicePrefixProperties{
				Prefix: to.StringPtr("<prefix>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ServicePrefixesClientCreateOrUpdateResult)
}
```
