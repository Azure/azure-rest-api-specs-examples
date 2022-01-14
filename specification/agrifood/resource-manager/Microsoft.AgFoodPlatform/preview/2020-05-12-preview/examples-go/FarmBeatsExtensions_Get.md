Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fagrifood%2Farmagrifood%2Fv0.3.0/sdk/resourcemanager/agrifood/armagrifood/README.md) on how to add the SDK to your project and authenticate.

```go
package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2020-05-12-preview/examples/FarmBeatsExtensions_Get.json
func ExampleFarmBeatsExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armagrifood.NewFarmBeatsExtensionsClient(cred, nil)
	res, err := client.Get(ctx,
		"<farm-beats-extension-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FarmBeatsExtensionsClientGetResult)
}
```
