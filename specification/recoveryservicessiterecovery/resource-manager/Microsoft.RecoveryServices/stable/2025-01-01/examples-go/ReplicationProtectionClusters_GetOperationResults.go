package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionClusters_GetOperationResults.json
func ExampleReplicationProtectionClustersClient_GetOperationResults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationProtectionClustersClient().GetOperationResults(ctx, "resourceGroupPS1", "vault1", "eastus", "eastus-container", "cluster1", "ea63a935-59d5-4b12-aff2-98773f63c116", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ReplicationProtectionCluster = armrecoveryservicessiterecovery.ReplicationProtectionCluster{
	// 	Name: to.Ptr("cluster12"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionClusters"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/eastus/replicationProtectionContainers/eastus-container/replicationProtectionClusters/cluster12"),
	// 	Properties: &armrecoveryservicessiterecovery.ReplicationProtectionClusterProperties{
	// 		ActiveLocation: to.Ptr("Primary"),
	// 		AllowedOperations: []*string{
	// 			to.Ptr("UnplannedFailover"),
	// 			to.Ptr("DisableProtection")},
	// 			AreAllClusterNodesRegistered: to.Ptr(false),
	// 			ClusterNodeFqdns: []*string{
	// 			},
	// 			ClusterProtectedItemIDs: []*string{
	// 			},
	// 			ClusterRegisteredNodes: []*armrecoveryservicessiterecovery.RegisteredClusterNodes{
	// 			},
	// 			HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 			},
	// 			PolicyFriendlyName: to.Ptr("24-hour-retention-policy"),
	// 			PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/24-hour-retention-policy"),
	// 			PrimaryFabricFriendlyName: to.Ptr("eastus"),
	// 			PrimaryFabricProvider: to.Ptr("AzureFabric"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("eastus"),
	// 			ProtectionState: to.Ptr("NotProtected"),
	// 			ProtectionStateDescription: to.Ptr("NotProtected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationProtectionClusterDetails{
	// 				InstanceType: to.Ptr("A2A"),
	// 				ClusterManagementID: to.Ptr("4b8c1c01-5ad7-4958-a98d-daf7cca6a124"),
	// 				InitialPrimaryFabricLocation: to.Ptr("eastus"),
	// 				InitialPrimaryZone: to.Ptr(""),
	// 				InitialRecoveryFabricLocation: to.Ptr("centraluseuap"),
	// 				InitialRecoveryZone: to.Ptr(""),
	// 				LifecycleID: to.Ptr("d5f6f011-7ef5-4c78-80fe-6e1b28401c4d"),
	// 				PrimaryFabricLocation: to.Ptr("eastus"),
	// 				RecoveryFabricLocation: to.Ptr("centraluseuap"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			RecoveryContainerID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/centraluseuap/replicationProtectionContainers/centraluseuap-container"),
	// 			RecoveryFabricFriendlyName: to.Ptr("centraluseuap"),
	// 			RecoveryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/centraluseuap"),
	// 			RecoveryProtectionContainerFriendlyName: to.Ptr("eastus-container"),
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			SharedDiskProperties: &armrecoveryservicessiterecovery.SharedDiskReplicationItemProperties{
	// 				AllowedOperations: []*string{
	// 				},
	// 				HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 				},
	// 				ProtectionState: to.Ptr("EnablingProtection"),
	// 				ReplicationHealth: to.Ptr("Normal"),
	// 				SharedDiskProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASharedDiskReplicationDetails{
	// 					InstanceType: to.Ptr("A2A"),
	// 					ManagementID: to.Ptr("4b8c1c01-5ad7-4958-a98d-daf7cca6a124"),
	// 					PrimaryFabricLocation: to.Ptr("eastus"),
	// 					ProtectedManagedDisks: []*armrecoveryservicessiterecovery.A2AProtectedManagedDiskDetails{
	// 					},
	// 					RecoveryFabricLocation: to.Ptr("centraluseuap"),
	// 				},
	// 			},
	// 		},
	// 	}
}
