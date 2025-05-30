package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ClusterRecoveryPoints_ListByReplicationProtectionCluster.json
func ExampleClusterRecoveryPointsClient_NewListByReplicationProtectionClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClusterRecoveryPointsClient().NewListByReplicationProtectionClusterPager("resourceGroupPS1", "vault1", "fabric-pri-eastus", "pri-cloud-eastus", "testcluster", nil)
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
		// page.ClusterRecoveryPointCollection = armrecoveryservicessiterecovery.ClusterRecoveryPointCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ClusterRecoveryPoint{
		// 		{
		// 			Name: to.Ptr("cc48b7f3-b267-432b-ad76-45528974dc62"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionClusters/recoveryPoints"),
		// 			ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster/recoveryPoints/cc48b7f3-b267-432b-ad76-45528974dc62"),
		// 			Properties: &armrecoveryservicessiterecovery.ClusterRecoveryPointProperties{
		// 				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AClusterRecoveryPointDetails{
		// 					InstanceType: to.Ptr("A2A"),
		// 					Nodes: []*string{
		// 						to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4"),
		// 						to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio")},
		// 						RecoveryPointSyncType: to.Ptr(armrecoveryservicessiterecovery.RecoveryPointSyncTypeMultiVMSyncRecoveryPoint),
		// 					},
		// 					RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-19T08:55:51.176Z"); return t}()),
		// 					RecoveryPointType: to.Ptr(armrecoveryservicessiterecovery.ClusterRecoveryPointTypeCrashConsistent),
		// 				},
		// 		}},
		// 	}
	}
}
