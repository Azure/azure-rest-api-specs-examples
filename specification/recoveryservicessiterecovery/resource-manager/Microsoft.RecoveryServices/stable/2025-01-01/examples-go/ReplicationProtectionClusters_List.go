package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionClusters_List.json
func ExampleReplicationProtectionClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationProtectionClustersClient().NewListPager("resourceGroupPS1", "vault1", &armrecoveryservicessiterecovery.ReplicationProtectionClustersClientListOptions{SkipToken: nil,
		Filter: to.Ptr("SourceFabricName eq 'asr-a2a-default-eastus' and SourceFabricLocation eq 'East US' and InstanceType eq 'A2A'"),
	})
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
		// page.ReplicationProtectionClusterCollection = armrecoveryservicessiterecovery.ReplicationProtectionClusterCollection{
		// 	Value: []*armrecoveryservicessiterecovery.ReplicationProtectionCluster{
		// 		{
		// 			Name: to.Ptr("cluster1"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionClusters"),
		// 			ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/eastus/replicationProtectionContainers/eastus-container/replicationProtectionClusters/cluster1"),
		// 			Properties: &armrecoveryservicessiterecovery.ReplicationProtectionClusterProperties{
		// 				ActiveLocation: to.Ptr("Primary"),
		// 				AgentClusterID: to.Ptr("dd9925cf-5f90-49e9-bd6d-d4cbdcec956b"),
		// 				AllowedOperations: []*string{
		// 					to.Ptr("UnplannedFailover"),
		// 					to.Ptr("DisableProtection")},
		// 					AreAllClusterNodesRegistered: to.Ptr(true),
		// 					ClusterFqdn: to.Ptr("ad45f2fc-f9d6-42ac-8a7c-1c5380c88c28"),
		// 					ClusterNodeFqdns: []*string{
		// 					},
		// 					ClusterProtectedItemIDs: []*string{
		// 						to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/eastus/replicationProtectionContainers/eastus-container/replicationProtectedItems/cluster2vm0"),
		// 						to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/eastus/replicationProtectionContainers/eastus-container/replicationProtectedItems/cluster2vm1")},
		// 						ClusterRegisteredNodes: []*armrecoveryservicessiterecovery.RegisteredClusterNodes{
		// 							{
		// 								BiosID: to.Ptr("eed7a457-b11f-420e-8e4f-384833a6b7c9"),
		// 								ClusterNodeFqdn: to.Ptr("VM0"),
		// 								IsSharedDiskVirtualNode: to.Ptr(true),
		// 								MachineID: to.Ptr("a7eaf02d-60d9-45be-a444-d1b945f1c7b2"),
		// 						}},
		// 						HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
		// 						},
		// 						PolicyFriendlyName: to.Ptr("24-hour-retention-policy"),
		// 						PolicyID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/24-hour-retention-policy"),
		// 						PrimaryFabricFriendlyName: to.Ptr("eastus"),
		// 						PrimaryFabricProvider: to.Ptr("AzureFabric"),
		// 						PrimaryProtectionContainerFriendlyName: to.Ptr("eastus"),
		// 						ProtectionClusterType: to.Ptr("WindowsServerFailoverCluster"),
		// 						ProtectionState: to.Ptr("Protected"),
		// 						ProtectionStateDescription: to.Ptr("Protected"),
		// 						ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationProtectionClusterDetails{
		// 							InstanceType: to.Ptr("A2A"),
		// 							ClusterManagementID: to.Ptr("1ed32804-ae51-4752-9448-9f686ae27d7b"),
		// 							InitialPrimaryFabricLocation: to.Ptr("eastus"),
		// 							InitialPrimaryZone: to.Ptr(""),
		// 							InitialRecoveryFabricLocation: to.Ptr("centraluseuap"),
		// 							InitialRecoveryZone: to.Ptr(""),
		// 							MultiVMGroupCreateOption: to.Ptr(armrecoveryservicessiterecovery.MultiVMGroupCreateOptionUserSpecified),
		// 							MultiVMGroupID: to.Ptr("4c2988ed-7e2a-566e-9c5f-c33835621e83"),
		// 							MultiVMGroupName: to.Ptr("multiVmGroupName1111"),
		// 							PrimaryFabricLocation: to.Ptr("eastus"),
		// 							RecoveryFabricLocation: to.Ptr("centraluseuap"),
		// 						},
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						RecoveryContainerID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/centraluseuap/replicationProtectionContainers/centraluseuap-container"),
		// 						RecoveryFabricFriendlyName: to.Ptr("centraluseuap"),
		// 						RecoveryFabricID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/centraluseuap"),
		// 						RecoveryProtectionContainerFriendlyName: to.Ptr("eastus-container"),
		// 						ReplicationHealth: to.Ptr("Normal"),
		// 						SharedDiskProperties: &armrecoveryservicessiterecovery.SharedDiskReplicationItemProperties{
		// 							AllowedOperations: []*string{
		// 							},
		// 							HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
		// 							},
		// 							ProtectionState: to.Ptr("Protected"),
		// 							ReplicationHealth: to.Ptr("Normal"),
		// 							SharedDiskProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASharedDiskReplicationDetails{
		// 								InstanceType: to.Ptr("A2A"),
		// 								LastRpoCalculatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-11T07:34:58.856Z"); return t}()),
		// 								ManagementID: to.Ptr("1ed32804-ae51-4752-9448-9f686ae27d7b"),
		// 								PrimaryFabricLocation: to.Ptr("eastus"),
		// 								RecoveryFabricLocation: to.Ptr("centraluseuap"),
		// 								RpoInSeconds: to.Ptr[int64](300),
		// 							},
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
