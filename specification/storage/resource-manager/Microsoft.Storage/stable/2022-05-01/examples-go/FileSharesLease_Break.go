package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storage/resource-manager/Microsoft.Storage/stable/2022-05-01/examples/FileSharesLease_Break.json
func ExampleFileSharesClient_Lease_fileSharesLease_Break() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorage.NewFileSharesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Lease(ctx, "res3376", "sto328", "share12", &armstorage.FileSharesClientLeaseOptions{XMSSnapshot: nil,
		Parameters: &armstorage.LeaseShareRequest{
			Action:  to.Ptr(armstorage.LeaseShareActionBreak),
			LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
