package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/PostPrivateStoreCollectionOffer.json
func ExamplePrivateStoreCollectionOfferClient_Post() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	_, err = client.Post(ctx,
		"<private-store-id>",
		"<offer-id>",
		"<collection-id>",
		&armmarketplace.PrivateStoreCollectionOfferClientPostOptions{Payload: nil})
	if err != nil {
		log.Fatal(err)
	}
}
