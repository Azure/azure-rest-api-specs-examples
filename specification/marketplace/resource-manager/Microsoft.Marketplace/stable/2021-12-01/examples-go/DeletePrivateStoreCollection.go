package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-12-01/examples/DeletePrivateStoreCollection.json
func ExamplePrivateStoreCollectionClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmarketplace.NewPrivateStoreCollectionClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"a0e28e55-90c4-41d8-8e34-bb7ef7775406",
		"d0f5aa2c-ecc3-4d87-906a-f8c486dcc4f1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
