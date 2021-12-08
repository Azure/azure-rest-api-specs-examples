Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdigitaltwins%2Farmdigitaltwins%2Fv0.1.0/sdk/resourcemanager/digitaltwins/armdigitaltwins/README.md) on how to add the SDK to your project and authenticate.

```go
package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2020-12-01/examples/DigitalTwinsEndpointsGet_example.json
func ExampleDigitalTwinsEndpointClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdigitaltwins.NewDigitalTwinsEndpointClient("<subscription-id>", cred, nil)
	pager := client.List("<resource-group-name>",
		"<resource-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DigitalTwinsEndpointResource.ID: %s\n", *v.ID)
		}
	}
}
```
