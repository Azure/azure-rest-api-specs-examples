package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2026-03-01-preview/Clusters_List_MaximumSet_Gen.json
func ExampleClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("1E4BD993-6890-4E69-8966-81482D7502EF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClustersClient().NewListPager("rgopenapi", "myOrganization", "myProject", nil)
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
		// page = armmongodbatlas.ClustersClientListResponse{
		// 	ClusterListResult: armmongodbatlas.ClusterListResult{
		// 		Value: []*armmongodbatlas.Cluster{
		// 			{
		// 				Properties: &armmongodbatlas.ClusterProperties{
		// 					ClusterName: to.Ptr("myCluster"),
		// 					ClusterTier: to.Ptr(armmongodbatlas.ClusterTierFREE),
		// 					RegionName: to.Ptr("eastus"),
		// 					MongoDbVersion: to.Ptr("7.0"),
		// 					Backups: to.Ptr(true),
		// 					State: to.Ptr("IDLE"),
		// 					ProvisioningState: to.Ptr(armmongodbatlas.ResourceProvisioningStateSucceeded),
		// 				},
		// 				ID: to.Ptr("/subscriptions/BC19EA3E-EDFC-4B6E-875E-B220CCDE497D/resourceGroups/rgopenapi/providers/MongoDB.Atlas/organizations/myOrganization/projects/myProject/clusters/myCluster"),
		// 				Name: to.Ptr("myCluster"),
		// 				Type: to.Ptr("MongoDB.Atlas/organizations/projects/clusters"),
		// 				SystemData: &armmongodbatlas.SystemData{
		// 					CreatedBy: to.Ptr("rxmwavptbhnfjbl"),
		// 					CreatedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-18T10:44:11.098Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("vvyjmiqdxlzkubu"),
		// 					LastModifiedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-03-18T10:44:11.099Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
