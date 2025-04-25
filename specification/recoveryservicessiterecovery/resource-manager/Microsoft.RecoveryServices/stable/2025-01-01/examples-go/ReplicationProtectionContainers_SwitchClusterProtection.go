package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionContainers_SwitchClusterProtection.json
func ExampleReplicationProtectionContainersClient_BeginSwitchClusterProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectionContainersClient().BeginSwitchClusterProtection(ctx, "resourceGroupPS1", "vault1", "fabric-pri-eastus", "pri-cloud-eastus", armrecoveryservicessiterecovery.SwitchClusterProtectionInput{
		Properties: &armrecoveryservicessiterecovery.SwitchClusterProtectionInputProperties{
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASwitchClusterProtectionInput{
				InstanceType: to.Ptr("A2A"),
				PolicyID:     to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/klncksan"),
				ProtectedItemsDetail: []*armrecoveryservicessiterecovery.A2AProtectedItemDetail{
					{
						RecoveryResourceGroupID:      to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
						ReplicationProtectedItemName: to.Ptr("yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio"),
						VMManagedDisks: []*armrecoveryservicessiterecovery.A2AVMManagedDiskInputDetails{
							{
								DiskID:                              to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql0-osdisk"),
								PrimaryStagingAzureStorageAccountID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache"),
								RecoveryResourceGroupID:             to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
							}},
					},
					{
						RecoveryResourceGroupID:      to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
						ReplicationProtectedItemName: to.Ptr("kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4"),
						VMManagedDisks: []*armrecoveryservicessiterecovery.A2AVMManagedDiskInputDetails{
							{
								DiskID:                              to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql1-osdisk"),
								PrimaryStagingAzureStorageAccountID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/clustertestrg-19-01/providers/Microsoft.Storage/storageAccounts/ix701lvaasrcache"),
								RecoveryResourceGroupID:             to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
							}},
					}},
				RecoveryContainerID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-rec-westus/replicationProtectionContainers/rec-cloud-westus"),
			},
			ReplicationProtectionClusterName: to.Ptr("testcluster"),
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
	// res.ProtectionContainer = armrecoveryservicessiterecovery.ProtectionContainer{
	// 	Name: to.Ptr("pri-cloud-eastus"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus"),
	// 	Properties: &armrecoveryservicessiterecovery.ProtectionContainerProperties{
	// 		FabricFriendlyName: to.Ptr("East US"),
	// 		FabricType: to.Ptr("Azure"),
	// 		FriendlyName: to.Ptr("pri-cloud-eastus"),
	// 		PairingStatus: to.Ptr("Paired"),
	// 		ProtectedItemCount: to.Ptr[int32](13),
	// 		Role: to.Ptr("Primary"),
	// 	},
	// }
}
