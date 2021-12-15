Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmfeatures%2Fv0.1.1/sdk/resourcemanager/resources/armfeatures/README.md) on how to add the SDK to your project and authenticate.

```go
package armfeatures_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/unregisterFeature.json
func ExampleFeaturesClient_Unregister() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfeatures.NewFeaturesClient("<subscription-id>", cred, nil)
	res, err := client.Unregister(ctx,
		"<resource-provider-namespace>",
		"<feature-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("FeatureResult.ID: %s\n", *res.ID)
}
```
