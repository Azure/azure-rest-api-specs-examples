package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/Clusters_ListBySubscription.json
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
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster"),
		// 				Name: to.Ptr("examplecluster"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters"),
		// 				Location: to.Ptr("westus2"),
		// 				Tags: map[string]*string{
		// 					"env": to.Ptr("dev"),
		// 				},
		// 				Properties: &armhorizondb.ClusterProperties{
		// 					AdministratorLogin: to.Ptr("exampleadministratorlogin"),
		// 					Version: to.Ptr("17"),
		// 					VCores: to.Ptr[int32](4),
		// 					ReplicaCount: to.Ptr[int32](2),
		// 					ZonePlacementPolicy: to.Ptr(armhorizondb.ZonePlacementPolicyStrict),
		// 					State: to.Ptr(armhorizondb.StateReady),
		// 					FullyQualifiedDomainName: to.Ptr("examplecluster.westus2.horizondb.azure.com"),
		// 					ReadonlyEndpoint: to.Ptr("examplecluster-ro.westus2.horizondb.azure.com"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster2"),
		// 				Name: to.Ptr("examplecluster2"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters"),
		// 				Location: to.Ptr("westus2"),
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("development"),
		// 				},
		// 				Properties: &armhorizondb.ClusterProperties{
		// 					AdministratorLogin: to.Ptr("admin"),
		// 					Version: to.Ptr("17"),
		// 					VCores: to.Ptr[int32](2),
		// 					ReplicaCount: to.Ptr[int32](1),
		// 					ZonePlacementPolicy: to.Ptr(armhorizondb.ZonePlacementPolicyStrict),
		// 					State: to.Ptr(armhorizondb.StateReady),
		// 					FullyQualifiedDomainName: to.Ptr("examplecluster2.westus2.horizondb.azure.com"),
		// 					ReadonlyEndpoint: to.Ptr("examplecluster2-ro.westus2.horizondb.azure.com"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
