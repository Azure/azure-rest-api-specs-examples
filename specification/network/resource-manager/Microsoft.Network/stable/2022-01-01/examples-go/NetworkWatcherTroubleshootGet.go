package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/NetworkWatcherTroubleshootGet.json
func ExampleWatchersClient_BeginGetTroubleshooting() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewWatchersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginGetTroubleshooting(ctx, "rg1", "nw1", armnetwork.TroubleshootingParameters{
		Properties: &armnetwork.TroubleshootingProperties{
			StorageID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/st1"),
			StoragePath: to.Ptr("https://st1.blob.core.windows.net/cn1"),
		},
		TargetResourceID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
