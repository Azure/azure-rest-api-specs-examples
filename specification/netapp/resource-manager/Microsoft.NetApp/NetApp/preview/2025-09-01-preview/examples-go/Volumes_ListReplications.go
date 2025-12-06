package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/Volumes_ListReplications.json
func ExampleVolumesClient_NewListReplicationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListReplicationsPager("myRG", "account1", "pool1", "volume1", &armnetapp.VolumesClientListReplicationsOptions{
		Body: &armnetapp.ListReplicationsRequest{
			Exclude: to.Ptr(armnetapp.ExcludeNone),
		}})
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
		// page = armnetapp.VolumesClientListReplicationsResponse{
		// 	ListReplications: armnetapp.ListReplications{
		// 		Value: []*armnetapp.Replication{
		// 			{
		// 				ReplicationSchedule: to.Ptr(armnetapp.ReplicationScheduleDaily),
		// 				RemoteVolumeResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-5999/capacityPools/pool-0977/volumes/volume-4508"),
		// 				RemoteVolumeRegion: to.Ptr("westus"),
		// 				MirrorState: to.Ptr(armnetapp.ReplicationMirrorStateMirrored),
		// 				ReplicationCreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-15T13:23:33Z"); return t}()),
		// 				ReplicationDeletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-08-16T13:23:33Z"); return t}()),
		// 			},
		// 		},
		// 	},
		// }
	}
}
