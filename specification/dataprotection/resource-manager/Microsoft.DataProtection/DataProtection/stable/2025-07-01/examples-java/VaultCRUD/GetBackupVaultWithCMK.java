
/**
 * Samples for BackupVaults GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VaultCRUD/GetBackupVaultWithCMK.json
     */
    /**
     * Sample code: Get BackupVault With CMK.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupVaultWithCMK(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
