
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AUpdateReplicationProtectedItemInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.A2AVmManagedDiskUpdateDetails;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.ReplicationProtectedItem;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateConfidentialDiskEncryptionInfo;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.UpdateReplicationProtectedItemInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationProtectedItems Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationProtectedItems_Update_A2ACvm.json
     */
    /**
     * Sample code: Updates the target confidential disk encryption set for an A2A confidential VM.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void updatesTheTargetConfidentialDiskEncryptionSetForAnA2AConfidentialVM(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        ReplicationProtectedItem resource
            = manager.replicationProtectedItems().getWithResponse("priyanprg", "priyanponeboxvault", "CentralUSCanSite",
                "CentralUSCancloud", "a2aCvmVm", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new UpdateReplicationProtectedItemInputProperties()
            .withProviderSpecificDetails(new A2AUpdateReplicationProtectedItemInput()
                .withRecoveryResourceGroupId(
                    "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg")
                .withManagedDiskUpdateDetails(Arrays.asList(new A2AVmManagedDiskUpdateDetails().withDiskId(
                    "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/disks/a2aCvmVm-osdisk")
                    .withRecoveryTargetDiskAccountType("Premium_LRS").withRecoveryReplicaDiskAccountType("Premium_LRS")
                    .withConfidentialDiskEncryptionInfo(
                        new UpdateConfidentialDiskEncryptionInfo().withRecoveryTargetConfidentialDiskEncryptionSetId(
                            "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.Compute/diskEncryptionSets/cvm-target-des"))))
                .withRecoveryConfidentialDataDiskEncryptionIdentity(
                    "/Subscriptions/42195872-7e70-4f8a-837f-84b28ecbb78b/resourceGroups/priyanprg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/cvm-cdde-uami")))
            .apply();
    }
}
