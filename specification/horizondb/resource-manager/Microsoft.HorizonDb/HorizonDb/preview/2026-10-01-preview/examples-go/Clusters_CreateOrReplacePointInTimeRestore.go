package armhorizondb_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Clusters_CreateOrReplacePointInTimeRestore.json
func ExampleClustersClient_BeginCreateOrReplace_restoreAHorizonDbClusterToAPointInTime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrReplace(ctx, "example-resource-group", "restored-cluster", armhorizondb.Cluster{
		Location: to.Ptr("eastus2"),
		Properties: &armhorizondb.ClusterProperties{
			CreationMode:                    to.Ptr(armhorizondb.ClusterCreateModePointInTimeRestore),
			PointInTimeUTC:                  to.Ptr(time.Date(2026, time.July, 27, 12, 0, 0, 0, time.UTC)),
			CreationSourceClusterResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/source-cluster"),
			Network: &armhorizondb.Network{
				PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
			},
			Backup: &armhorizondb.BackupProperties{
				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateEnabled),
				RetentionDays: to.Ptr[int32](14),
			},
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
	// res = armhorizondb.ClustersClientCreateOrReplaceResponse{
	// 	Cluster: armhorizondb.Cluster{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/restored-cluster"),
	// 		Name: to.Ptr("restored-cluster"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters"),
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armhorizondb.ClusterProperties{
	// 			PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV17),
	// 			CreationMode: to.Ptr(armhorizondb.ClusterCreateModePointInTimeRestore),
	// 			PointInTimeUTC: to.Ptr(time.Date(2026, time.July, 27, 12, 0, 0, 0, time.UTC)),
	// 			CreationSourceClusterResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/source-cluster"),
	// 			Network: &armhorizondb.Network{
	// 				PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
	// 			},
	// 			Backup: &armhorizondb.BackupProperties{
	// 				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateEnabled),
	// 				RetentionDays: to.Ptr[int32](14),
	// 			},
	// 			State: to.Ptr(armhorizondb.ClusterStateProvisioning),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateProvisioning),
	// 			ReadWritePool: &armhorizondb.ReadWritePoolProperties{
	// 			},
	// 		},
	// 	},
	// }
}
