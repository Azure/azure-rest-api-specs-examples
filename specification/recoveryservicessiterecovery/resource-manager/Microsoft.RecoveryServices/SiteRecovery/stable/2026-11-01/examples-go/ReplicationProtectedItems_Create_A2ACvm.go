package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v4"
)

// Generated from example definition: 2026-11-01/ReplicationProtectedItems_Create_A2ACvm.json
func ExampleReplicationProtectedItemsClient_BeginCreate_enablesProtectionForAnA2AConfidentialVMWithCmkDiskEncryption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("42195872-7e70-4f8a-837f-84b28ecbb78b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectedItemsClient().BeginCreate(ctx, "priyanprg", "priyanponeboxvault", "CentralUSCanSite", "CentralUSCancloud", "a2aCvmVm", armrecoveryservicessiterecovery.EnableProtectionInput{
		Properties: &armrecoveryservicessiterecovery.EnableProtectionInputProperties{
			PolicyID:          to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationPolicies/a2aPolicy"),
			ProtectableItemID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectableItems/a2aCvmVm"),
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AEnableProtectionInput{
				InstanceType:            to.Ptr("A2A"),
				FabricObjectID:          to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/virtualMachines/a2aCvmVm"),
				RecoveryContainerID:     to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/EUSCanSite/replicationProtectionContainers/euscancloud"),
				RecoveryResourceGroupID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg"),
				VMManagedDisks: []*armrecoveryservicessiterecovery.A2AVMManagedDiskInputDetails{
					{
						DiskID:                              to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aCvmVm-osdisk"),
						PrimaryStagingAzureStorageAccountID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Storage/storageAccounts/a2acache"),
						RecoveryResourceGroupID:             to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg"),
						RecoveryReplicaDiskAccountType:      to.Ptr("Premium_LRS"),
						RecoveryTargetDiskAccountType:       to.Ptr("Premium_LRS"),
						ConfidentialDiskEncryptionInfo: &armrecoveryservicessiterecovery.ConfidentialDiskEncryptionInfo{
							RecoveryReplicaConfidentialDiskEncryptionSetID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-replica-des"),
							RecoveryTargetConfidentialDiskEncryptionSetID:  to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-target-des"),
						},
					},
				},
				RecoveryConfidentialDataDiskEncryptionIdentity: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/cvm-cdde-uami"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicessiterecovery.ReplicationProtectedItemsClientCreateResponse{
	// 	ReplicationProtectedItem: armrecoveryservicessiterecovery.ReplicationProtectedItem{
	// 		Name: to.Ptr("a2aCvmVm"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers/replicationProtectedItems"),
	// 		ID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectedItems/a2aCvmVm"),
	// 		Properties: &armrecoveryservicessiterecovery.ReplicationProtectedItemProperties{
	// 			ActiveLocation: to.Ptr("Primary"),
	// 			AllowedOperations: []*string{
	// 				to.Ptr("PlannedFailover"),
	// 				to.Ptr("UnplannedFailover"),
	// 				to.Ptr("DisableProtection"),
	// 				to.Ptr("TestFailover"),
	// 			},
	// 			FriendlyName: to.Ptr("a2aCvmVm"),
	// 			PolicyFriendlyName: to.Ptr("a2aPolicy"),
	// 			PolicyID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationPolicies/a2aPolicy"),
	// 			PrimaryFabricFriendlyName: to.Ptr("Central US"),
	// 			PrimaryProtectionContainerFriendlyName: to.Ptr("CentralUSCancloud"),
	// 			ProtectableItemID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectableItems/a2aCvmVm"),
	// 			ProtectedItemType: to.Ptr("A2A"),
	// 			ProtectionState: to.Ptr("Protected"),
	// 			ProtectionStateDescription: to.Ptr("Protected"),
	// 			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2AReplicationDetails{
	// 				InstanceType: to.Ptr("A2A"),
	// 				FabricObjectID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/virtualMachines/a2aCvmVm"),
	// 				RecoveryConfidentialDataDiskEncryptionIdentity: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/cvm-cdde-uami"),
	// 				ProtectedManagedDisks: []*armrecoveryservicessiterecovery.A2AProtectedManagedDiskDetails{
	// 					{
	// 						DiskID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aCvmVm-osdisk"),
	// 						RecoveryResourceGroupID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg"),
	// 						RecoveryReplicaDiskAccountType: to.Ptr("Premium_LRS"),
	// 						RecoveryTargetDiskAccountType: to.Ptr("Premium_LRS"),
	// 						ConfidentialDiskEncryptionInfo: &armrecoveryservicessiterecovery.ConfidentialDiskEncryptionInfo{
	// 							RecoveryReplicaConfidentialDiskEncryptionSetID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-replica-des"),
	// 							RecoveryTargetConfidentialDiskEncryptionSetID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-target-des"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ReplicationHealth: to.Ptr("Normal"),
	// 			TestFailoverState: to.Ptr("None"),
	// 			TestFailoverStateDescription: to.Ptr("None"),
	// 		},
	// 	},
	// }
}
