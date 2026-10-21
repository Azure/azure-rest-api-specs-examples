package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-10-01-preview/Clusters_Update.json
func ExampleClustersClient_BeginUpdate_updateAHorizonDbCluster() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "example-resource-group", "example-cluster", armhorizondb.ClusterForPatchUpdate{
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
			"team":        to.Ptr("data-platform"),
		},
		Properties: &armhorizondb.ClusterPropertiesForPatchUpdate{
			Network: &armhorizondb.NetworkForPatchUpdate{
				PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
			},
			Backup: &armhorizondb.BackupPropertiesForPatchUpdate{
				RetentionDays: to.Ptr[int32](21),
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
	// res = armhorizondb.ClustersClientUpdateResponse{
	// 	Cluster: armhorizondb.Cluster{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster"),
	// 		Name: to.Ptr("example-cluster"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters"),
	// 		Location: to.Ptr("eastus2"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"team": to.Ptr("data-platform"),
	// 		},
	// 		Identity: &armhorizondb.ClusterManagedServiceIdentity{
	// 			Type: to.Ptr(armhorizondb.ClusterIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armhorizondb.UserAssignedIdentity{
	// 				"/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/fabric-writer": &armhorizondb.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 					ClientID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 				},
	// 			},
	// 		},
	// 		Properties: &armhorizondb.ClusterProperties{
	// 			PostgreSQLVersion: to.Ptr(armhorizondb.PostgreSQLVersionV17),
	// 			CreationMode: to.Ptr(armhorizondb.ClusterCreateModeCreate),
	// 			Network: &armhorizondb.Network{
	// 				PublicAccess: to.Ptr(armhorizondb.PublicNetworkAccessStateDisabled),
	// 			},
	// 			Backup: &armhorizondb.BackupProperties{
	// 				GeoRedundancy: to.Ptr(armhorizondb.GeoRedundancyStateEnabled),
	// 				RetentionDays: to.Ptr[int32](21),
	// 				EarliestRestorePointUTC: to.Ptr(time.Date(2026, time.July, 27, 2, 0, 0, 0, time.UTC)),
	// 			},
	// 			State: to.Ptr(armhorizondb.ClusterStateReady),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 			ReadWritePool: &armhorizondb.ReadWritePoolProperties{
	// 				DesignatedPoolResourceID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/example-resource-group/providers/Microsoft.HorizonDb/clusters/example-cluster/pools/example-pool-01"),
	// 			},
	// 		},
	// 	},
	// }
}
