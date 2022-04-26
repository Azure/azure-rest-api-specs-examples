Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplace%2Farmmarketplace%2Fv0.4.0/sdk/resourcemanager/marketplace/armmarketplace/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/TransferOffers.json
func ExamplePrivateStoreCollectionClient_TransferOffers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.TransferOffers(ctx,
		"<private-store-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionClientTransferOffersOptions{Payload: &armmarketplace.TransferOffersProperties{
			Properties: &armmarketplace.TransferOffersDetails{
				OfferIDsList: []*string{
					to.Ptr("marketplacetestthirdparty.md-test-third-party-2"),
					to.Ptr("marketplacetestthirdparty.md-test-third-party-3")},
				Operation: to.Ptr("<operation>"),
				TargetCollections: []*string{
					to.Ptr("c752f021-1c37-4af5-b82f-74c51c27b44a"),
					to.Ptr("f47ef1c7-e908-4f39-ae29-db181634ad8d")},
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
