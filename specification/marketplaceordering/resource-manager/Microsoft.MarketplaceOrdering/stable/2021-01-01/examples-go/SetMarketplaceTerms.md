Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplaceordering%2Farmmarketplaceordering%2Fv0.4.0/sdk/resourcemanager/marketplaceordering/armmarketplaceordering/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplaceordering_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplaceordering/armmarketplaceordering"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SetMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Create(ctx,
		armmarketplaceordering.OfferTypeVirtualmachine,
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		armmarketplaceordering.AgreementTerms{
			Properties: &armmarketplaceordering.AgreementProperties{
				Accepted:             to.Ptr(false),
				LicenseTextLink:      to.Ptr("<license-text-link>"),
				MarketplaceTermsLink: to.Ptr("<marketplace-terms-link>"),
				Plan:                 to.Ptr("<plan>"),
				PrivacyPolicyLink:    to.Ptr("<privacy-policy-link>"),
				Product:              to.Ptr("<product>"),
				Publisher:            to.Ptr("<publisher>"),
				RetrieveDatetime:     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.12132Z"); return t }()),
				Signature:            to.Ptr("<signature>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
