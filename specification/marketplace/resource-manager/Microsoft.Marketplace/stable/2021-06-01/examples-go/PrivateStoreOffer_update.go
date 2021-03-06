package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/PrivateStoreOffer_update.json
func ExamplePrivateStoreCollectionOfferClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<private-store-id>",
		"<offer-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionOfferClientCreateOrUpdateOptions{Payload: &armmarketplace.Offer{
			Properties: &armmarketplace.OfferProperties{
				ETag: to.StringPtr("<etag>"),
				SpecificPlanIDsLimitation: []*string{
					to.StringPtr("0001"),
					to.StringPtr("0002")},
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateStoreCollectionOfferClientCreateOrUpdateResult)
}
