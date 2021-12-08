Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaps%2Farmmaps%2Fv0.1.0/sdk/resourcemanager/maps/armmaps/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/CreateMapsCreator.json
func ExampleCreatorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaps.NewCreatorsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<creator-name>",
		armmaps.Creator{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Creator.ID: %s\n", *res.ID)
}
```
