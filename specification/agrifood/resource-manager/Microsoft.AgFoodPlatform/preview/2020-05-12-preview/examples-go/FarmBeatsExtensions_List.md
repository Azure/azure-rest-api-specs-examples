Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fagrifood%2Farmagrifood%2Fv0.5.0/sdk/resourcemanager/agrifood/armagrifood/README.md) on how to add the SDK to your project and authenticate.

```go
package armagrifood_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2020-05-12-preview/examples/FarmBeatsExtensions_List.json
func ExampleFarmBeatsExtensionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armagrifood.NewFarmBeatsExtensionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListPager(&armagrifood.FarmBeatsExtensionsClientListOptions{FarmBeatsExtensionIDs: []string{},
		FarmBeatsExtensionNames: []string{},
		ExtensionCategories:     []string{},
		PublisherIDs:            []string{},
		MaxPageSize:             nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
