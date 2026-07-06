
/**
 * Samples for StorageAccounts Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountFailover.json
     */
    /**
     * Sample code: StorageAccountFailover.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountFailover(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().failover("res4228", "sto2434", null,
            com.azure.core.util.Context.NONE);
    }
}
