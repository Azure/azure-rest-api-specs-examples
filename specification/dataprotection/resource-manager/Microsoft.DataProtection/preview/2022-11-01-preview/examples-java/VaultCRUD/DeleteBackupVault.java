/** Samples for BackupVaults Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/VaultCRUD/DeleteBackupVault.json
     */
    /**
     * Sample code: Delete BackupVault.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void deleteBackupVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupVaults()
            .deleteByResourceGroupWithResponse(
                "SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
