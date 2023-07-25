package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Volumes_ListReplications.json
func ExampleVolumesClient_NewListReplicationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumesClient().NewListReplicationsPager("myRG", "account1", "pool1", "volume1", nil)
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
		// page.ListReplications = armnetapp.ListReplications{
		// 	Value: []*armnetapp.Replication{
		// 		{
		// 			RemoteVolumeRegion: to.Ptr("westus"),
		// 			RemoteVolumeResourceID: to.Ptr("/subscriptions/36e85c76-e720-473e-881f-e2fe72f462d0/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account-5999/capacityPools/pool-0977/volumes/volume-4508"),
		// 			ReplicationSchedule: to.Ptr(armnetapp.ReplicationScheduleDaily),
		// 	}},
		// }
	}
}
