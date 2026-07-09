package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2026-03-01-preview/Clusters_CreateOrUpdate_MaximumSet_Gen.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("1E4BD993-6890-4E69-8966-81482D7502EF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClustersClient().BeginCreateOrUpdate(ctx, "rgopenapi", "myOrganization", "myProject", "myCluster", armmongodbatlas.Cluster{
		Properties: &armmongodbatlas.ClusterProperties{
			ClusterTier: to.Ptr(armmongodbatlas.ClusterTierFREE),
			RegionName:  to.Ptr("eastus"),
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
	// res = armmongodbatlas.ClustersClientCreateOrUpdateResponse{
	// 	Cluster: armmongodbatlas.Cluster{
	// 		ID: to.Ptr("/subscriptions/1E4BD993-6890-4E69-8966-81482D7502EF/resourceGroups/rgopenapi/providers/MongoDB.Atlas/organizations/myOrganization/projects/myProject/clusters/myCluster"),
	// 		Name: to.Ptr("myCluster"),
	// 		Type: to.Ptr("MongoDB.Atlas/organizations/projects/clusters"),
	// 		Properties: &armmongodbatlas.ClusterProperties{
	// 			ClusterName: to.Ptr("myCluster"),
	// 			ClusterTier: to.Ptr(armmongodbatlas.ClusterTierFREE),
	// 			RegionName: to.Ptr("eastus"),
	// 			ProvisioningState: to.Ptr(armmongodbatlas.ResourceProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armmongodbatlas.SystemData{
	// 			CreatedBy: to.Ptr("rxmwavptbhnfjbl"),
	// 			CreatedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-18T10:44:11.098Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("vvyjmiqdxlzkubu"),
	// 			LastModifiedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-18T10:44:11.099Z"); return t}()),
	// 		},
	// 	},
	// }
}
