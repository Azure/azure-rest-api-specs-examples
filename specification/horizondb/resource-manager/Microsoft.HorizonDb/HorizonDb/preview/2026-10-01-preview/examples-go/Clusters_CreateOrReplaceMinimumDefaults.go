package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Clusters_CreateOrReplaceMinimumDefaults.json
func ExampleClustersClient_BeginCreateOrReplace_createAPostgreSql18HorizonDbClusterWithMinimumConfigurationAndServiceDefaults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrReplace(ctx, "example-resource-group", "minimum-defaults", armhorizondb.Cluster{
		Location: to.Ptr("eastus2"),
		Properties: &armhorizondb.ClusterProperties{
			PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV18),
			Backup: &armhorizondb.BackupProperties{
				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateDisabled),
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
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/minimum-defaults"),
	// 		Name: to.Ptr("minimum-defaults"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters"),
	// 		Location: to.Ptr("eastus2"),
	// 		Properties: &armhorizondb.ClusterProperties{
	// 			PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV18),
	// 			CreationMode: to.Ptr(armhorizondb.ClusterCreateModeCreate),
	// 			Network: &armhorizondb.Network{
	// 				PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
	// 			},
	// 			Backup: &armhorizondb.BackupProperties{
	// 				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateDisabled),
	// 				RetentionDays: to.Ptr[int32](7),
	// 			},
	// 			State: to.Ptr(armhorizondb.ClusterStateUpdating),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateInProgress),
	// 			ReadWritePool: &armhorizondb.ReadWritePoolProperties{
	// 			},
	// 		},
	// 	},
	// }
}
