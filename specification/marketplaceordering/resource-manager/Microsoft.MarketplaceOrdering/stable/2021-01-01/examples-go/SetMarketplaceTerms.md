Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplaceordering%2Farmmarketplaceordering%2Fv0.2.1/sdk/resourcemanager/marketplaceordering/armmarketplaceordering/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/SetMarketplaceTerms.json
func ExampleMarketplaceAgreementsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplaceordering.NewMarketplaceAgreementsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		armmarketplaceordering.OfferType("virtualmachine"),
		"<publisher-id>",
		"<offer-id>",
		"<plan-id>",
		armmarketplaceordering.AgreementTerms{
			Properties: &armmarketplaceordering.AgreementProperties{
				Accepted:             to.BoolPtr(false),
				LicenseTextLink:      to.StringPtr("<license-text-link>"),
				MarketplaceTermsLink: to.StringPtr("<marketplace-terms-link>"),
				Plan:                 to.StringPtr("<plan>"),
				PrivacyPolicyLink:    to.StringPtr("<privacy-policy-link>"),
				Product:              to.StringPtr("<product>"),
				Publisher:            to.StringPtr("<publisher>"),
				RetrieveDatetime:     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T11:33:07.12132Z"); return t }()),
				Signature:            to.StringPtr("<signature>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.MarketplaceAgreementsClientCreateResult)
}
```
