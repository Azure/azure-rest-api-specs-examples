Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdigitaltwins%2Farmdigitaltwins%2Fv0.1.0/sdk/resourcemanager/digitaltwins/armdigitaltwins/README.md) on how to add the SDK to your project and authenticate.

```go
package armdigitaltwins_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
)

// x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2020-12-01/examples/DigitalTwinsEndpointGet_example.json
func ExampleDigitalTwinsEndpointClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdigitaltwins.NewDigitalTwinsEndpointClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<endpoint-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DigitalTwinsEndpointResource.ID: %s\n", *res.ID)
}
```
