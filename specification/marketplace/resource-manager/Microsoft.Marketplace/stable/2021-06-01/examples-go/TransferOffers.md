```go
package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/TransferOffers.json
func ExamplePrivateStoreCollectionClient_TransferOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	res, err := client.TransferOffers(ctx,
		"<private-store-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionClientTransferOffersOptions{Payload: &armmarketplace.TransferOffersProperties{
			Properties: &armmarketplace.TransferOffersDetails{
				OfferIDsList: []*string{
					to.StringPtr("marketplacetestthirdparty.md-test-third-party-2"),
					to.StringPtr("marketplacetestthirdparty.md-test-third-party-3")},
				Operation: to.StringPtr("<operation>"),
				TargetCollections: []*string{
					to.StringPtr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
					to.StringPtr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateStoreCollectionClientTransferOffersResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplace%2Farmmarketplace%2Fv0.2.1/sdk/resourcemanager/marketplace/armmarketplace/README.md) on how to add the SDK to your project and authenticate.
