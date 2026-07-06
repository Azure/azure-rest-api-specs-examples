
/**
 * Samples for StorageAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountDelete.json
     */
    /**
     * Sample code: StorageAccountDelete.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountDelete(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().deleteWithResponse("res4228", "sto2434",
            com.azure.core.util.Context.NONE);
    }
}
