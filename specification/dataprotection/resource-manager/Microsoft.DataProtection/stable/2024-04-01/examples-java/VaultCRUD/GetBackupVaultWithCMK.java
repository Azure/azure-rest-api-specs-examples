
/**
 * Samples for BackupVaults GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/
     * GetBackupVaultWithCMK.json
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
