/** Samples for BackupVaults GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/VaultCRUD/GetBackupVaultWithMSI.json
     */
    /**
     * Sample code: Get BackupVault With MSI.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupVaultWithMSI(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupVaults()
            .getByResourceGroupWithResponse("SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
