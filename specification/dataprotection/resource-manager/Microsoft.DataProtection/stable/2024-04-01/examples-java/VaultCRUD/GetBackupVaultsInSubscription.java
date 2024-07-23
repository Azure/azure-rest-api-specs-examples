
/**
 * Samples for BackupVaults List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/VaultCRUD/
     * GetBackupVaultsInSubscription.json
     */
    /**
     * Sample code: Get BackupVaults in Subscription.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getBackupVaultsInSubscription(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().list(com.azure.core.util.Context.NONE);
    }
}
