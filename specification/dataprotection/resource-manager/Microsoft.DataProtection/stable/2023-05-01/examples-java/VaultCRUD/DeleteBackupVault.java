/** Samples for BackupVaults Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/VaultCRUD/DeleteBackupVault.json
     */
    /**
     * Sample code: Delete BackupVault.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteBackupVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().delete("SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
