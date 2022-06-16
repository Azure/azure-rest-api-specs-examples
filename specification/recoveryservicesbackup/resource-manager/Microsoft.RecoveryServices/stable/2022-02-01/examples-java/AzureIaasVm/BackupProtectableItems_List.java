import com.azure.core.util.Context;

/** Samples for BackupProtectableItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureIaasVm/BackupProtectableItems_List.json
     */
    /**
     * Sample code: List protectable items with backupManagementType filter as AzureIaasVm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectableItemsWithBackupManagementTypeFilterAsAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupProtectableItems()
            .list("NetSDKTestRsVault", "SwaggerTestRg", "backupManagementType eq 'AzureIaasVM'", null, Context.NONE);
    }
}
