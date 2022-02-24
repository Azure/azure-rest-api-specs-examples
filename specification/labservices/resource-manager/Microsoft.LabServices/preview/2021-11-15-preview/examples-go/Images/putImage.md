Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Flabservices%2Farmlabservices%2Fv0.2.1/sdk/resourcemanager/labservices/armlabservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armlabservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
)

// x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Images/putImage.json
func ExampleImagesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlabservices.NewImagesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<lab-plan-name>",
		"<image-name>",
		armlabservices.Image{
			Properties: &armlabservices.ImageProperties{
				EnabledState: armlabservices.EnableStateEnabled.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ImagesClientCreateOrUpdateResult)
}
```
