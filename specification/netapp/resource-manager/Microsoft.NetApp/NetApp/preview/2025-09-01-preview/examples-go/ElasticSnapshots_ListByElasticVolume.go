package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v8"
)

// Generated from example definition: 2025-09-01-preview/ElasticSnapshots_ListByElasticVolume.json
func ExampleElasticSnapshotsClient_NewListByElasticVolumePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewElasticSnapshotsClient().NewListByElasticVolumePager("myRG", "account1", "pool1", "volume1", nil)
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
		// page = armnetapp.ElasticSnapshotsClientListByElasticVolumeResponse{
		// 	ElasticSnapshotListResult: armnetapp.ElasticSnapshotListResult{
		// 		Value: []*armnetapp.ElasticSnapshot{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/elasticAccounts/account1/elasticCapacityPools/pool1/elasticVolumes/volume1/elasticSnapshots/snapshot1"),
		// 				Name: to.Ptr("account1/pool1/volume1/snapshot1"),
		// 				Type: to.Ptr("Microsoft.NetApp/elasticAccounts/elasticCapacityPools/elasticVolumes/elasticSnapshots"),
		// 				Properties: &armnetapp.ElasticSnapshotProperties{
		// 					ProvisioningState: to.Ptr(armnetapp.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
