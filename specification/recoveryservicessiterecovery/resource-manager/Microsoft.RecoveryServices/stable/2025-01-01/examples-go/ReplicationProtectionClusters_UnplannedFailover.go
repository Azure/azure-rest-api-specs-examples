package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationProtectionClusters_UnplannedFailover.json
func ExampleReplicationProtectionClustersClient_BeginUnplannedFailover() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectionClustersClient().BeginUnplannedFailover(ctx, "resourceGroupPS1", "vault1", "fabric-pri-eastus", "pri-cloud-eastus", "testcluster", armrecoveryservicessiterecovery.ClusterUnplannedFailoverInput{
		Properties: &armrecoveryservicessiterecovery.ClusterUnplannedFailoverInputProperties{
			FailoverDirection: to.Ptr("primarytorecovery"),
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AClusterUnplannedFailoverInput{
				InstanceType:           to.Ptr("A2A"),
				ClusterRecoveryPointID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster/recoveryPoints/cc48b7f3-b267-432b-ad76-45528974dc62"),
				IndividualNodeRecoveryPoints: []*string{
					to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/testVM/recoveryPoints/b5c2051b-79e3-41ad-9d25-244f6ef8ce7d")},
			},
			SourceSiteOperations: to.Ptr("NotRequired"),
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
	// res.ReplicationProtectionCluster = armrecoveryservicessiterecovery.ReplicationProtectionCluster{
	// 	Name: to.Ptr("testcluster"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectionClusters"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster"),
	// 	Properties: &armrecoveryservicessiterecovery.ReplicationProtectionClusterProperties{
	// 		ActiveLocation: to.Ptr("Recovery"),
	// 		AgentClusterID: to.Ptr("aeeca6a3-171b-4c9d-ae22-0e4adb6416a0"),
	// 		AllowedOperations: []*string{
	// 			to.Ptr("CommitFailoverProtectionCluster"),
	// 			to.Ptr("ChangePitProtectionCluster"),
	// 			to.Ptr("PurgeProtectionCluster")},
	// 			AreAllClusterNodesRegistered: to.Ptr(true),
	// 			ClusterFqdn: to.Ptr("sdgqlc"),
	// 			ClusterNodeFqdns: []*string{
	// 				to.Ptr("sdgql0"),
	// 				to.Ptr("sdgql1")},
	// 				ClusterProtectedItemIDs: []*string{
	// 					to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/yNdYnDYKZ7hYU7zyVeBychFBCyAbEkrJcJNUarDrXio"),
	// 					to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectedItems/kdUdWvpVnm3QgOQPHoVMX8YAtAO8OC4kKNjt40ERSr4")},
	// 					ClusterRegisteredNodes: []*armrecoveryservicessiterecovery.RegisteredClusterNodes{
	// 						{
	// 							BiosID: to.Ptr("37D0059C-9118-4220-AA1D-58A10EFA7660"),
	// 							ClusterNodeFqdn: to.Ptr("sdgql0"),
	// 							IsSharedDiskVirtualNode: to.Ptr(false),
	// 							MachineID: to.Ptr("3794026b-6792-4d12-9c0c-de0c79376c90"),
	// 						},
	// 						{
	// 							BiosID: to.Ptr("aeeca6a3-171b-4c9d-ae22-0e4adb6416a0"),
	// 							ClusterNodeFqdn: to.Ptr("aeeca6a3-171b-4c9d-ae22-0e4adb6416a0"),
	// 							IsSharedDiskVirtualNode: to.Ptr(true),
	// 							MachineID: to.Ptr("aeeca6a3-171b-4c9d-ae22-0e4adb6416a0"),
	// 						},
	// 						{
	// 							BiosID: to.Ptr("83CCE932-67EC-4C13-AB29-ACF5F8F7ED48"),
	// 							ClusterNodeFqdn: to.Ptr("sdgql1"),
	// 							IsSharedDiskVirtualNode: to.Ptr(false),
	// 							MachineID: to.Ptr("cb27913a-a5f2-4691-9eba-78b67f45a57a"),
	// 					}},
	// 					CurrentScenario: &armrecoveryservicessiterecovery.CurrentScenarioDetails{
	// 						JobID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationJobs/None"),
	// 						ScenarioName: to.Ptr("None"),
	// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1753-01-01T01:01:01.000Z"); return t}()),
	// 					},
	// 					HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 					},
	// 					LastSuccessfulFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-22T09:31:13.160Z"); return t}()),
	// 					LastSuccessfulTestFailoverTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-22T06:08:01.744Z"); return t}()),
	// 					PolicyFriendlyName: to.Ptr("klncksan"),
	// 					PolicyID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationPolicies/klncksan"),
	// 					PrimaryFabricFriendlyName: to.Ptr("East US"),
	// 					PrimaryFabricProvider: to.Ptr("AzureFabric"),
	// 					PrimaryProtectionContainerFriendlyName: to.Ptr("East US"),
	// 					ProtectionClusterType: to.Ptr("WindowsServerFailoverCluster"),
	// 					ProtectionState: to.Ptr("FailoverCompleted"),
	// 					ProtectionStateDescription: to.Ptr("Failover completed."),
	// 					ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationProtectionClusterDetails{
	// 						InstanceType: to.Ptr("A2A"),
	// 						ClusterManagementID: to.Ptr("a24d47b2-a80b-4553-9c16-499c30c3be07"),
	// 						FailoverRecoveryPointID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster/recoveryPoints/cc48b7f3-b267-432b-ad76-45528974dc62"),
	// 						InitialPrimaryFabricLocation: to.Ptr("eastus"),
	// 						InitialPrimaryZone: to.Ptr(""),
	// 						InitialRecoveryFabricLocation: to.Ptr("westus"),
	// 						InitialRecoveryZone: to.Ptr(""),
	// 						LastRpoCalculatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-22T10:19:11.049Z"); return t}()),
	// 						LifecycleID: to.Ptr("3d523ab9-8c6d-40af-850d-ccee06513dc6"),
	// 						MultiVMGroupCreateOption: to.Ptr(armrecoveryservicessiterecovery.MultiVMGroupCreateOptionUserSpecified),
	// 						MultiVMGroupID: to.Ptr("a7ef77cb-ae59-545f-a32a-bf30575ab1c6"),
	// 						MultiVMGroupName: to.Ptr("testcluster"),
	// 						PrimaryFabricLocation: to.Ptr("eastus"),
	// 						RecoveryFabricLocation: to.Ptr("westus"),
	// 						RpoInSeconds: to.Ptr[int64](3213),
	// 					},
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					RecoveryContainerID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-rec-westus/replicationProtectionContainers/rec-cloud-westus"),
	// 					RecoveryFabricFriendlyName: to.Ptr("West US"),
	// 					RecoveryFabricID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-rec-westus"),
	// 					RecoveryProtectionContainerFriendlyName: to.Ptr("pri-cloud-eastus"),
	// 					ReplicationHealth: to.Ptr("Normal"),
	// 					SharedDiskProperties: &armrecoveryservicessiterecovery.SharedDiskReplicationItemProperties{
	// 						AllowedOperations: []*string{
	// 						},
	// 						HealthErrors: []*armrecoveryservicessiterecovery.HealthError{
	// 						},
	// 						ProtectionState: to.Ptr("UnplannedFailoverCommitRequired"),
	// 						ReplicationHealth: to.Ptr("Normal"),
	// 						SharedDiskProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASharedDiskReplicationDetails{
	// 							InstanceType: to.Ptr("A2A"),
	// 							FailoverRecoveryPointID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/fabric-pri-eastus/replicationProtectionContainers/pri-cloud-eastus/replicationProtectionClusters/testcluster/recoveryPoints/cc48b7f3-b267-432b-ad76-45528974dc62"),
	// 							LastRpoCalculatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-22T09:20:35.930Z"); return t}()),
	// 							ManagementID: to.Ptr("a24d47b2-a80b-4553-9c16-499c30c3be07"),
	// 							PrimaryFabricLocation: to.Ptr("eastus"),
	// 							ProtectedManagedDisks: []*armrecoveryservicessiterecovery.A2AProtectedManagedDiskDetails{
	// 								{
	// 									AllowedDiskLevelOperation: []*string{
	// 									},
	// 									DataPendingAtSourceAgentInMB: to.Ptr[float64](0),
	// 									DataPendingInStagingStorageAccountInMB: to.Ptr[float64](0),
	// 									DiskCapacityInBytes: to.Ptr[int64](274877906944),
	// 									DiskID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourcegroups/clustertestrg-19-01/providers/microsoft.compute/disks/sdgql-datadisk0"),
	// 									DiskName: to.Ptr("sdgql-datadisk0"),
	// 									DiskState: to.Ptr("Protected"),
	// 									DiskType: to.Ptr("Data"),
	// 									FailoverDiskName: to.Ptr("sdgql-datadisk0"),
	// 									IsDiskEncrypted: to.Ptr(false),
	// 									IsDiskKeyEncrypted: to.Ptr(false),
	// 									PrimaryStagingAzureStorageAccountID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/resourceGroupPS1/providers/Microsoft.Storage/storageAccounts/ix701lshashankvaasrcache"),
	// 									RecoveryReplicaDiskAccountType: to.Ptr("Premium_LRS"),
	// 									RecoveryReplicaDiskID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr/providers/Microsoft.Compute/disks/sdgql-datadisk0-ASRReplica"),
	// 									RecoveryResourceGroupID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr"),
	// 									RecoveryTargetDiskAccountType: to.Ptr("Premium_LRS"),
	// 									RecoveryTargetDiskID: to.Ptr("/subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/ClusterTestRG-19-01-asr/providers/Microsoft.Compute/disks/sdgql-datadisk0"),
	// 									ResyncRequired: to.Ptr(false),
	// 									TfoDiskName: to.Ptr("sdgql-datadisk0-ASRtest"),
	// 							}},
	// 							RecoveryFabricLocation: to.Ptr("westus"),
	// 							RpoInSeconds: to.Ptr[int64](10),
	// 						},
	// 						TestFailoverState: to.Ptr("None"),
	// 					},
	// 					TestFailoverState: to.Ptr("None"),
	// 					TestFailoverStateDescription: to.Ptr("None"),
	// 				},
	// 			}
}
