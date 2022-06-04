Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv1.0.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemList.json
func ExampleAnalyticsItemsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewAnalyticsItemsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"my-resource-group",
		"my-component",
		armapplicationinsights.ItemScopePathAnalyticsItems,
		&armapplicationinsights.AnalyticsItemsClientListOptions{Scope: nil,
			Type:           nil,
			IncludeContent: nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
