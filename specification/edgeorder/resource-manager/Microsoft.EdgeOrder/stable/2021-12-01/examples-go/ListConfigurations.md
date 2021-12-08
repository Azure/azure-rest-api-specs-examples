Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fedgeorder%2Farmedgeorder%2Fv0.1.0/sdk/resourcemanager/edgeorder/armedgeorder/README.md) on how to add the SDK to your project and authenticate.

```go
package armedgeorder_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"
)

// x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListConfigurations.json
func ExampleEdgeOrderManagementClient_ListConfigurations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armedgeorder.NewEdgeOrderManagementClient("<subscription-id>", cred, nil)
	pager := client.ListConfigurations(armedgeorder.ConfigurationsRequest{
		ConfigurationFilters: []*armedgeorder.ConfigurationFilters{
			{
				FilterableProperty: []*armedgeorder.FilterableProperty{
					{
						Type: armedgeorder.SupportedFilterTypesShipToCountries.ToPtr(),
						SupportedValues: []*string{
							to.StringPtr("US")},
					}},
				HierarchyInformation: &armedgeorder.HierarchyInformation{
					ProductFamilyName: to.StringPtr("<product-family-name>"),
					ProductLineName:   to.StringPtr("<product-line-name>"),
					ProductName:       to.StringPtr("<product-name>"),
				},
			}},
	},
		&armedgeorder.EdgeOrderManagementClientListConfigurationsOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
```
