import com.azure.core.util.Context;

/** Samples for BackupProtectedItems List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureIaasVm/BackupProtectedItems_List.json
     */
    /**
     * Sample code: List protected items with backupManagementType filter as AzureIaasVm.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectedItemsWithBackupManagementTypeFilterAsAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .backupProtectedItems()
            .list(
                "NetSDKTestRsVault",
                "SwaggerTestRg",
                "backupManagementType eq 'AzureIaasVM' and itemType eq 'VM'",
                null,
                Context.NONE);
    }
}
