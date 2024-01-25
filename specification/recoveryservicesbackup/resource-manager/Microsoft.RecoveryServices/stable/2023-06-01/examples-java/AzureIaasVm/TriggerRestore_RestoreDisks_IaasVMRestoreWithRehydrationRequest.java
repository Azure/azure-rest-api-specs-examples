
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreWithRehydrationRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryPointRehydrationInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RehydrationPriority;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;

/**
 * Samples for Restores Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/
     * AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreWithRehydrationRequest.json
     */
    /**
     * Sample code: Restore Disks with IaasVMRestoreWithRehydrationRequest.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreDisksWithIaasVMRestoreWithRehydrationRequest(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.restores().trigger("testVault", "netsdktestrg", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334",
            new RestoreRequestResource().withProperties(new IaasVMRestoreWithRehydrationRequest()
                .withRecoveryPointId("348916168024334").withRecoveryType(RecoveryType.RESTORE_DISKS)
                .withSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                .withStorageAccountId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount")
                .withRegion("southeastasia").withCreateNewCloudService(true).withOriginalStorageAccountOption(false)
                .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                .withRecoveryPointRehydrationInfo(new RecoveryPointRehydrationInfo()
                    .withRehydrationRetentionDuration("P7D").withRehydrationPriority(RehydrationPriority.STANDARD))),
            com.azure.core.util.Context.NONE);
    }
}
