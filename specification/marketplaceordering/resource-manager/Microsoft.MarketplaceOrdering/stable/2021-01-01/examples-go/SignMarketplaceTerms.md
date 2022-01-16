Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplaceordering%2Farmmarketplaceordering%2Fv0.2.0/sdk/resourcemanager/marketplaceordering/armmarketplaceordering/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplaceordering_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
)

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SignMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Sign() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Sign(ctx,
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientSignResult)
}
```
