package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ClusterRecoveryPoint_Get.json
func ExampleClusterRecoveryPointClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClusterRecoveryPointClient().Get(ctx, "resourceGroupPS1", "vault1", "fabric-pri-eastus", "pri-cloud-eastus", "testcluster", "06b9ae7f-f21d-4a76-9897-5cf5d6004d80", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterRecoveryPoint = armrecoveryservicessiterecovery.ClusterRecoveryPoint{
	// 	Name: to.Ptr("06b9ae7f-f21d-4a76-9897-5cf5d6004d80"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionClusters/recoveryPoints"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster/recoveryPoints/06b9ae7f-f21d-4a76-9897-5cf5d6004d80"),
	// 	Properties: &armrecoveryservicessiterecovery.ClusterRecoveryPointProperties{
	// 		ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AClusterRecoveryPointDetails{
	// 			InstanceType: to.Ptr("A2A"),
	// 			Nodes: []*string{
	// 				to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4"),
	// 				to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio")},
	// 				RecoveryPointSyncType: to.Ptr(armrecoveryservicessiterecovery.RecoveryPointSyncTypeMultiVMSyncRecoveryPoint),
	// 			},
	// 			RecoveryPointTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-19T10:51:29.632Z"); return t}()),
	// 			RecoveryPointType: to.Ptr(armrecoveryservicessiterecovery.ClusterRecoveryPointTypeCrashConsistent),
	// 		},
	// 	}
}
