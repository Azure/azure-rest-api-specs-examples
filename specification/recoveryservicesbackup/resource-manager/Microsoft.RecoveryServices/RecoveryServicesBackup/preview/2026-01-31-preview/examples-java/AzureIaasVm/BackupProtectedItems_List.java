
/**
 * Samples for BackupProtectedItems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureIaasVm/BackupProtectedItems_List.json
     */
    /**
     * Sample code: List protected items with backupManagementType filter as AzureIaasVm.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void listProtectedItemsWithBackupManagementTypeFilterAsAzureIaasVm(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.backupProtectedItems().list("NetSDKTestRsVault", "SwaggerTestRg",
            "backupManagementType eq 'AzureIaasVM' and itemType eq 'VM'", null, com.azure.core.util.Context.NONE);
    }
}
