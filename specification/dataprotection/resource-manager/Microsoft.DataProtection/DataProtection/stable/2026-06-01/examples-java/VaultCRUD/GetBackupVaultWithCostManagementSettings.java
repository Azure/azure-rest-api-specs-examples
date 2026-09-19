
/**
 * Samples for BackupVaults GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/VaultCRUD/GetBackupVaultWithCostManagementSettings.json
     */
    /**
     * Sample code: Get BackupVault with Cost Management Settings.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupVaultWithCostManagementSettings(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
