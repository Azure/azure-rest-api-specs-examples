
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AEnableProtectionInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AVmManagedDiskInputDetails;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ConfidentialDiskEncryptionInfo;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.EnableProtectionInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationProtectedItems Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Create_A2ACvm.json
     */
    /**
     * Sample code: Enables protection for an A2A confidential VM with CMK disk encryption.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void enablesProtectionForAnA2AConfidentialVMWithCMKDiskEncryption(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationProtectedItems().define("a2aCvmVm")
            .withExistingReplicationProtectionContainer("priyanprg", "priyanponeboxvault", "CentralUSCanSite",
                "CentralUSCancloud")
            .withProperties(new EnableProtectionInputProperties().withPolicyId(
                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationPolicies/a2aPolicy")
                .withProtectableItemId(
                    "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/CentralUSCanSite/replicationProtectionContainers/CentralUSCancloud/replicationProtectableItems/a2aCvmVm")
                .withProviderSpecificDetails(new A2AEnableProtectionInput().withFabricObjectId(
                    "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/virtualMachines/a2aCvmVm")
                    .withRecoveryContainerId(
                        "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.RecoveryServices/vaults/priyanponeboxvault/replicationFabrics/EUSCanSite/replicationProtectionContainers/euscancloud")
                    .withRecoveryResourceGroupId(
                        "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg")
                    .withVmManagedDisks(Arrays.asList(new A2AVmManagedDiskInputDetails().withDiskId(
                        "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aCvmVm-osdisk")
                        .withPrimaryStagingAzureStorageAccountId(
                            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Storage/storageAccounts/a2acache")
                        .withRecoveryResourceGroupId(
                            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg")
                        .withRecoveryReplicaDiskAccountType("Premium_LRS")
                        .withRecoveryTargetDiskAccountType("Premium_LRS")
                        .withConfidentialDiskEncryptionInfo(new ConfidentialDiskEncryptionInfo()
                            .withRecoveryReplicaConfidentialDiskEncryptionSetId(
                                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-replica-des")
                            .withRecoveryTargetConfidentialDiskEncryptionSetId(
                                "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-target-des"))))
                    .withRecoveryConfidentialDataDiskEncryptionIdentity(
                        "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/cvm-cdde-uami")))
            .create();
    }
}
