package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/DeletePrivateStoreOffer.json
func ExamplePrivateStoreCollectionOfferClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionOfferClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"marketplacetestthirdparty.md-test-third-party-2",
		"56a1a02d-8cf8-45df-bf37-d5f7120fcb3d",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
