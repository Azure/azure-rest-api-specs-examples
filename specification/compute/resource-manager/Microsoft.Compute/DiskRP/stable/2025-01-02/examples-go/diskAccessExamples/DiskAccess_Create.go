package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskAccessExamples/DiskAccess_Create.json
func ExampleDiskAccessesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskAccessesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskAccess", armcompute.DiskAccess{
		Location: to.Ptr("West US"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskAccess = armcompute.DiskAccess{
	// 	Name: to.Ptr("myDiskAccess"),
	// 	Type: to.Ptr("Microsoft.Compute/diskAccesses"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourcegroup/providers/Microsoft.Compute/diskAccesses/myDiskAccess"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.DiskAccessProperties{
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-05-01T04:41:35.079Z"); return t}()),
	// 	},
	// }
}
