package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/Replicas_List.json
func ExampleReplicasClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicasClient().NewListPager("exampleresourcegroup", "examplecluster", "examplepool", nil)
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
		// page = armhorizondb.ReplicasClientListResponse{
		// 	ReplicaListResult: armhorizondb.ReplicaListResult{
		// 		Value: []*armhorizondb.Replica{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/pools/examplepool/replicas/examplereplica1"),
		// 				Name: to.Ptr("examplereplica1"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/replicas"),
		// 				Properties: &armhorizondb.ReplicaProperties{
		// 					Role: to.Ptr(armhorizondb.ReplicaRoleRead),
		// 					Status: to.Ptr(armhorizondb.StateReady),
		// 					FullyQualifiedDomainName: to.Ptr("examplereplica1.westus2.horizondb.azure.com"),
		// 					AvailabilityZone: to.Ptr("1"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/pools/examplepool/replicas/examplereplica2"),
		// 				Name: to.Ptr("examplereplica2"),
		// 				Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/replicas"),
		// 				Properties: &armhorizondb.ReplicaProperties{
		// 					Role: to.Ptr(armhorizondb.ReplicaRoleReadWrite),
		// 					Status: to.Ptr(armhorizondb.StateReady),
		// 					FullyQualifiedDomainName: to.Ptr("examplereplica2.westus2.horizondb.azure.com"),
		// 					AvailabilityZone: to.Ptr("2"),
		// 					ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
