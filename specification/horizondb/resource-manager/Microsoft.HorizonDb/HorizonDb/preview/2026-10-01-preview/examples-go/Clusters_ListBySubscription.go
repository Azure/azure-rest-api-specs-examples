package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Clusters_ListBySubscription.json
func ExampleClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListBySubscriptionPager(nil)
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
		// page = armhorizondb.ClustersClientListBySubscriptionResponse{
		// 	ClusterListResult: armhorizondb.ClusterListResult{
		// 		Value: []*armhorizondb.Cluster{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster"),
		// 				Name: to.Ptr("example-cluster"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters"),
		// 				Location: to.Ptr("eastus2"),
		// 				Properties: &armhorizondb.ClusterProperties{
		// 					PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV17),
		// 					CreationMode: to.Ptr(armhorizondb.ClusterCreateModeCreate),
		// 					Network: &armhorizondb.Network{
		// 						PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
		// 					},
		// 					Backup: &armhorizondb.BackupProperties{
		// 						GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateEnabled),
		// 						RetentionDays: to.Ptr[int32](14),
		// 						EarliestRestorePointUTC: to.Ptr(time.Date(2026, time.July, 27, 2, 0, 0, 0, time.UTC)),
		// 					},
		// 					State: to.Ptr(armhorizondb.ClusterStateReady),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 					ReadWritePool: &armhorizondb.ReadWritePoolProperties{
		// 						DesignatedPoolResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
