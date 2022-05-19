Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplaceordering%2Farmmarketplaceordering%2Fv1.0.0/sdk/resourcemanager/marketplaceordering/armmarketplaceordering/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplaceordering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SignMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Sign() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplaceordering.NewMarketplaceAgreementsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Sign(ctx,
		"pubid",
		"offid",
		"planid",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
