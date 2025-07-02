package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/86c6306649b02e542117adb46c61e8019dbd78e9/specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/BlobContainersLease_Break.json
func ExampleBlobContainersClient_Lease_breakALeaseOnAContainer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBlobContainersClient().Lease(ctx, "res3376", "sto328", "container6185", &armstorage.BlobContainersClientLeaseOptions{Parameters: &armstorage.LeaseContainerRequest{
		Action:  to.Ptr(armstorage.LeaseContainerRequestActionBreak),
		LeaseID: to.Ptr("8698f513-fa75-44a1-b8eb-30ba336af27d"),
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LeaseContainerResponse = armstorage.LeaseContainerResponse{
	// 	LeaseTimeSeconds: to.Ptr("0"),
	// }
}
