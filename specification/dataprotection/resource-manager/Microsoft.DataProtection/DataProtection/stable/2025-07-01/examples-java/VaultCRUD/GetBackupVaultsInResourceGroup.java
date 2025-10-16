
/**
 * Samples for BackupVaults ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VaultCRUD/GetBackupVaultsInResourceGroup.json
     */
    /**
     * Sample code: Get BackupVaults in ResourceGroup.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getBackupVaultsInResourceGroup(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().listByResourceGroup("SampleResourceGroup", com.azure.core.util.Context.NONE);
    }
}
