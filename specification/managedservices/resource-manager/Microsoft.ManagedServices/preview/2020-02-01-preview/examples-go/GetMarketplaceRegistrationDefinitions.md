Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagedservices%2Farmmanagedservices%2Fv0.1.0/sdk/resourcemanager/managedservices/armmanagedservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// x-ms-original-file: specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2020-02-01-preview/examples/GetMarketplaceRegistrationDefinitions.json
func ExampleMarketplaceRegistrationDefinitionsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedservices.NewMarketplaceRegistrationDefinitionsClient(cred, nil)
	pager := client.List("<scope>",
		&armmanagedservices.MarketplaceRegistrationDefinitionsListOptions{Filter: to.StringPtr("<filter>")})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("MarketplaceRegistrationDefinition.ID: %s\n", *v.ID)
		}
	}
}
```
