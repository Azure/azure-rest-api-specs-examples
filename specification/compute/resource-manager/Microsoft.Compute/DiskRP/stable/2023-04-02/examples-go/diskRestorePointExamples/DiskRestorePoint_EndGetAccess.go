package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/60679ee3db06e93eb73faa0587fed93ed843d6dc/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskRestorePointExamples/DiskRestorePoint_EndGetAccess.json
func ExampleDiskRestorePointClient_BeginRevokeAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskRestorePointClient().BeginRevokeAccess(ctx, "myResourceGroup", "rpc", "vmrp", "TestDisk45ceb03433006d1baee0_b70cd924-3362-4a80-93c2-9415eaa12745", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
