Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplace%2Farmmarketplace%2Fv0.2.1/sdk/resourcemanager/marketplace/armmarketplace/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/CreatePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<private-store-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionClientCreateOrUpdateOptions{Payload: &armmarketplace.Collection{
			Properties: &armmarketplace.CollectionProperties{
				AllSubscriptions: to.BoolPtr(false),
				Claim:            to.StringPtr("<claim>"),
				CollectionName:   to.StringPtr("<collection-name>"),
				SubscriptionsList: []*string{
					to.StringPtr("b340914e-353d-453a-85fb-8f9b65b51f91"),
					to.StringPtr("f2baa04d-5bfc-461b-b6d8-61b403c9ec48")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateStoreCollectionClientCreateOrUpdateResult)
}
```
