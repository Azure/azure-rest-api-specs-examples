package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v4"
)

// Generated from example definition: 2026-11-01/ReplicationProtectionContainers_SwitchProtection.json
func ExampleReplicationProtectionContainersClient_BeginSwitchProtection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("42195872-7e70-4f8a-837f-84b28ecbb78b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationProtectionContainersClient().BeginSwitchProtection(ctx, "priyanprg", "priyanponeboxvault", "CentralUSCanSite", "CentralUSCancloud", armrecoveryservicessiterecovery.SwitchProtectionInput{
		Properties: &armrecoveryservicessiterecovery.SwitchProtectionInputProperties{
			ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ASwitchProtectionInput{
				InstanceType:            to.Ptr("A2A"),
				RecoveryContainerID:     to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/EUSCanSite/replicationProtectionContainers/euscancloud"),
				RecoveryResourceGroupID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg"),
				VMManagedDisks: []*armrecoveryservicessiterecovery.A2AVMManagedDiskInputDetails{
					{
						DiskID:                              to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aSwapOsVm-osdisk"),
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
			ReplicationProtectedItemName: to.Ptr("a2aSwapOsVm"),
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
	// res = armrecoveryservicessiterecovery.ReplicationProtectionContainersClientSwitchProtectionResponse{
	// 	ProtectionContainer: armrecoveryservicessiterecovery.ProtectionContainer{
	// 		Name: to.Ptr("euscancloud"),
	// 		Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationProtectionContainers"),
	// 		ID: to.Ptr("/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/EUSCanSite/replicationProtectionContainers/euscancloud"),
	// 		Properties: &armrecoveryservicessiterecovery.ProtectionContainerProperties{
	// 			FabricFriendlyName: to.Ptr("East US 2 EUAP"),
	// 			FabricType: to.Ptr("Azure"),
	// 			FriendlyName: to.Ptr("euscancloud"),
	// 			PairingStatus: to.Ptr("Paired"),
	// 			ProtectedItemCount: to.Ptr[int32](0),
	// 			Role: to.Ptr("Primary"),
	// 		},
	// 	},
	// }
}
