package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/260ed6a52537921f53a18ffaf4020e3b4d510367/specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/FileSharesRestore.json
func ExampleFileSharesClient_Restore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFileSharesClient().Restore(ctx, "res3376", "sto328", "share1249", armstorage.DeletedShare{
		DeletedShareName:    to.Ptr("share1249"),
		DeletedShareVersion: to.Ptr("1234567890"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
