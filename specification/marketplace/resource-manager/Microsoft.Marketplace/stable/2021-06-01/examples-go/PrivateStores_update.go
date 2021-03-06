package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/PrivateStores_update.json
func ExamplePrivateStoreClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.CreateOrUpdate(ctx,
		"<private-store-id>",
		&armmarketplace.PrivateStoreClientCreateOrUpdateOptions{Payload: &armmarketplace.PrivateStore{
			Properties: &armmarketplace.PrivateStoreProperties{
				Availability: armmarketplace.Availability("disabled").ToPtr(),
				ETag:         to.StringPtr("<etag>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
