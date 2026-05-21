package armhorizondb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/horizondb/armhorizondb"
)

// Generated from example definition: 2026-01-20-preview/Replicas_Get.json
func ExampleReplicasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhorizondb.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicasClient().Get(ctx, "exampleresourcegroup", "examplecluster", "examplepool", "examplereplica", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhorizondb.ReplicasClientGetResponse{
	// 	Replica: &armhorizondb.Replica{
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.HorizonDb/clusters/examplecluster/pools/examplepool/replicas/examplereplica"),
	// 		Name: to.Ptr("examplereplica"),
	// 		Type: to.Ptr("Microsoft.HorizonDb/clusters/pools/replicas"),
	// 		Properties: &armhorizondb.ReplicaProperties{
	// 			Role: to.Ptr(armhorizondb.ReplicaRoleRead),
	// 			Status: to.Ptr(armhorizondb.StateHealthy),
	// 			FullyQualifiedDomainName: to.Ptr("examplereplica.westus2.horizondb.azure.com"),
	// 			AvailabilityZone: to.Ptr("1"),
	// 			ProvisioningState: to.Ptr(armhorizondb.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
