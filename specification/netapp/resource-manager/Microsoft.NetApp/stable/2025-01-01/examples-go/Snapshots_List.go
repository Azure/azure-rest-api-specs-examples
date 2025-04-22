package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/27046dbff974e3901970aa53b29cec6d8ec1342a/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-01-01/examples/Snapshots_List.json
func ExampleSnapshotsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSnapshotsClient().NewListPager("myRG", "account1", "pool1", "volume1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SnapshotsList = armnetapp.SnapshotsList{
		// 	Value: []*armnetapp.Snapshot{
		// 		{
		// 			Name: to.Ptr("account1/pool1/volume1/snapshot1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/snapshots"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/snapshots/snapshot1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.SnapshotProperties{
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33.000Z"); return t}()),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				SnapshotID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca3333"),
		// 			},
		// 	}},
		// }
	}
}
