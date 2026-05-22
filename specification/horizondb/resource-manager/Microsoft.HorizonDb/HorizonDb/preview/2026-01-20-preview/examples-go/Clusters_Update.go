package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/Clusters_Update.json
func ExampleClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginUpdate(ctx, "exampleresourcegroup", "examplecluster", armhorizondb.ClusterForPatchUpdate{
		Tags: map[string]*string{
			"team": to.Ptr("updated-data-platform"),
		},
		Properties: &armhorizondb.ClusterPropertiesForPatchUpdate{
			VCores:                     to.Ptr[int32](8),
			AdministratorLoginPassword: to.Ptr("examplenewpassword"),
		},
	}, nil)
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
	// res = armhorizondb.ClustersClientUpdateResponse{
	// 	Cluster: &armhorizondb.Cluster{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster"),
	// 		Name: to.Ptr("examplecluster"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters"),
	// 		Location: to.Ptr("westus2"),
	// 		Tags: map[string]*string{
	// 			"env": to.Ptr("dev"),
	// 			"team": to.Ptr("updated-data-platform"),
	// 		},
	// 		Properties: &armhorizondb.ClusterProperties{
	// 			AdministratorLogin: to.Ptr("exampleadministratorlogin"),
	// 			Version: to.Ptr("17"),
	// 			CreateMode: to.Ptr(armhorizondb.CreateModeClusterUpdate),
	// 			VCores: to.Ptr[int32](8),
	// 			ReplicaCount: to.Ptr[int32](2),
	// 			ZonePlacementPolicy: to.Ptr(armhorizondb.ZonePlacementPolicyStrict),
	// 			State: to.Ptr(armhorizondb.StateReady),
	// 			FullyQualifiedDomainName: to.Ptr("examplecluster.westus2.horizondb.azure.com"),
	// 			ReadonlyEndpoint: to.Ptr("examplecluster-ro.westus2.horizondb.azure.com"),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
