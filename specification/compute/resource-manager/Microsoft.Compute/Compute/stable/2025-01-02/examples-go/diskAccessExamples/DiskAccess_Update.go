package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/diskAccessExamples/DiskAccess_Update.json
func ExampleDiskAccessesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskAccessesClient().BeginUpdate(ctx, "myResourceGroup", "myDiskAccess", armcompute.DiskAccessUpdate{
		Tags: map[string]*string{
			"department": to.Ptr("Development"),
			"project":    to.Ptr("PrivateEndpoints"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DiskAccessesClientUpdateResponse{
	// 	DiskAccess: armcompute.DiskAccess{
	// 		Name: to.Ptr("myDiskAccess"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourcegroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess"),
	// 		Type: to.Ptr("Microsoft.Compute/diskAccesses"),
	// 		Location: to.Ptr("West US"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("Development"),
	// 			"project": to.Ptr("PrivateEndpoints"),
	// 		},
	// 	},
	// }
}
