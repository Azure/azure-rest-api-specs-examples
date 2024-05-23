
/**
 * Samples for BackupProtectedItems List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureIaasVm/BackupProtectedItems_List.json
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
