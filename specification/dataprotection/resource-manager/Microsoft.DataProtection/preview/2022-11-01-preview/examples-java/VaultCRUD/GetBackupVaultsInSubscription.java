/** Samples for BackupVaults List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/VaultCRUD/GetBackupVaultsInSubscription.json
     */
    /**
     * Sample code: Get BackupVaults in Subscription.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getBackupVaultsInSubscription(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().list(com.azure.core.util.Context.NONE);
    }
}
