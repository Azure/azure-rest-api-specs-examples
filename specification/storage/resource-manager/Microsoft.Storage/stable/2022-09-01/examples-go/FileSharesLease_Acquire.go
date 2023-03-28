package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b32e1896f30e6ea155449cb49719a6286e32b961/specification/storage/resource-manager/Microsoft.Storage/stable/2022-09-01/examples/FileSharesLease_Acquire.json
func ExampleFileSharesClient_Lease_acquireALeaseOnAShare() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSharesClient().Lease(ctx, "res3376", "sto328", "share124", &armstorage.FileSharesClientLeaseOptions{XMSSnapshot: nil,
		Parameters: &armstorage.LeaseShareRequest{
			Action:        to.Ptr(armstorage.LeaseShareActionAcquire),
			LeaseDuration: to.Ptr[int32](-1),
		},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LeaseShareResponse = armstorage.LeaseShareResponse{
	// 	LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
	// }
}
