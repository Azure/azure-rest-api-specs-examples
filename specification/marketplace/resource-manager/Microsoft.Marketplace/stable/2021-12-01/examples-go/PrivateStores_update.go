package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/PrivateStores_update.json
func ExamplePrivateStoreClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdate(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		&armmarketplace.PrivateStoreClientCreateOrUpdateOptions{Payload: &armmarketplace.PrivateStore{
			Properties: &armmarketplace.PrivateStoreProperties{
				Availability: to.Ptr(armmarketplace.AvailabilityDisabled),
				ETag:         to.Ptr("\"9301f4fd-0000-0100-0000-5e248b350345\""),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
