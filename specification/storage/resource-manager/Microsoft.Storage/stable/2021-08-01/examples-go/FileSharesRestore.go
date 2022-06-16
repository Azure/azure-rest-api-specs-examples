package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/FileSharesRestore.json
func ExampleFileSharesClient_Restore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewFileSharesClient("<subscription-id>", cred, nil)
	_, err = client.Restore(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<share-name>",
		armstorage.DeletedShare{
			DeletedShareName:    to.StringPtr("<deleted-share-name>"),
			DeletedShareVersion: to.StringPtr("<deleted-share-version>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
