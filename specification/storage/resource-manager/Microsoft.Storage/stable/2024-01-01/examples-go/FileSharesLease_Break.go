package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9d3547622288137fd36f086afcdaea5408dbe48c/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/FileSharesLease_Break.json
func ExampleFileSharesClient_Lease_breakALeaseOnAShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Lease(ctx, "res3376", "sto328", "share12", &armstorage.FileSharesClientLeaseOptions{XMSSnapshot: nil,
		Parameters: &armstorage.LeaseShareRequest{
			Action:  to.Ptr(armstorage.LeaseShareActionBreak),
			LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LeaseShareResponse = armstorage.LeaseShareResponse{
	// 	LeaseTimeSeconds: to.Ptr("0"),
	// }
}
