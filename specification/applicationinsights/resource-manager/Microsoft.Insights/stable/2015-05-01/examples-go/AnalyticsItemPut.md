Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fapplicationinsights%2Farmapplicationinsights%2Fv0.2.0/sdk/resourcemanager/applicationinsights/armapplicationinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnalyticsItemPut.json
func ExampleAnalyticsItemsClient_Put() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewAnalyticsItemsClient("<subscription-id>", cred, nil)
	res, err := client.Put(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.ItemScopePath("analyticsItems"),
		armapplicationinsights.ComponentAnalyticsItem{
			Content: to.StringPtr("<content>"),
			Name:    to.StringPtr("<name>"),
			Scope:   armapplicationinsights.ItemScope("shared").ToPtr(),
			Type:    armapplicationinsights.ItemType("query").ToPtr(),
		},
		&armapplicationinsights.AnalyticsItemsClientPutOptions{OverrideItem: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AnalyticsItemsClientPutResult)
}
```
