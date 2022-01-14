Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fagrifood%2Farmagrifood%2Fv0.3.0/sdk/resourcemanager/agrifood/armagrifood/README.md) on how to add the SDK to your project and authenticate.

```go
package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2020-05-12-preview/examples/FarmBeatsModels_Delete.json
func ExampleFarmBeatsModelsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armagrifood.NewFarmBeatsModelsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<farm-beats-resource-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
