package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5700885250d8f685a17293e930d98d1c1d72f401/specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/Snapshots_Update.json
func ExampleSnapshotsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSnapshotsClient().BeginUpdate(ctx, "myRG", "account1", "pool1", "volume1", "snapshot1", map[string]any{}, nil)
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
	// res.Snapshot = armnetapp.Snapshot{
	// 	Name: to.Ptr("account1/pool1/volume1/snapshot1"),
	// 	Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"),
	// 	ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetapp.SnapshotProperties{
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33.000Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SnapshotID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca3333"),
	// 	},
	// }
}
