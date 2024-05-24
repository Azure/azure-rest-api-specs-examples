
import com.azure.resourcemanager.recoveryservicesbackup.models.EncryptionDetails;
import com.azure.resourcemanager.recoveryservicesbackup.models.IaasVMRestoreRequest;
import com.azure.resourcemanager.recoveryservicesbackup.models.IdentityInfo;
import com.azure.resourcemanager.recoveryservicesbackup.models.RecoveryType;
import com.azure.resourcemanager.recoveryservicesbackup.models.RestoreRequestResource;
import com.azure.resourcemanager.recoveryservicesbackup.models.TargetDiskNetworkAccessOption;
import com.azure.resourcemanager.recoveryservicesbackup.models.TargetDiskNetworkAccessSettings;

/**
 * Samples for Restores Trigger.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureIaasVm/TriggerRestore_RestoreDisks_IaasVMRestoreRequest.json
     */
    /**
     * Sample code: Restore Disks with IaasVMRestoreRequest.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void restoreDisksWithIaasVMRestoreRequest(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.restores().trigger("testVault", "netsdktestrg", "Azure",
            "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
            "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334",
            new RestoreRequestResource().withProperties(new IaasVMRestoreRequest()
                .withRecoveryPointId("348916168024334").withRecoveryType(RecoveryType.RESTORE_DISKS)
                .withSourceResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1")
                .withStorageAccountId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testingRg/providers/Microsoft.Storage/storageAccounts/testAccount")
                .withRegion("southeastasia").withCreateNewCloudService(true).withOriginalStorageAccountOption(false)
                .withEncryptionDetails(new EncryptionDetails().withEncryptionEnabled(false))
                .withIdentityInfo(new IdentityInfo().withIsSystemAssignedIdentity(false).withManagedIdentityResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/asmaskarRG1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/asmaskartestmsi"))
                .withTargetDiskNetworkAccessSettings(new TargetDiskNetworkAccessSettings()
                    .withTargetDiskNetworkAccessOption(
                        TargetDiskNetworkAccessOption.ENABLE_PRIVATE_ACCESS_FOR_ALL_DISKS)
                    .withTargetDiskAccessId(
                        "/subscriptions/e7a191f5-713c-4bdb-b5e4-cf3dd90230ef/resourceGroups/arpja/providers/Microsoft.Compute/diskAccesses/arpja-diskaccess-ccy"))),
            com.azure.core.util.Context.NONE);
    }
}
