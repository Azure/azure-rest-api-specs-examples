package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9d3547622288137fd36f086afcdaea5408dbe48c/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobContainersDelete.json
func ExampleBlobContainersClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewBlobContainersClient().Delete(ctx, "res4079", "sto4506", "container9689", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
