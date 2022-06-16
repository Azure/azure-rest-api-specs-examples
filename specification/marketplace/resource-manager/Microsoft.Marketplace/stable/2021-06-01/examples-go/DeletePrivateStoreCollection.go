package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/DeletePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	_, err = client.Delete(ctx,
		"<private-store-id>",
		"<collection-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
