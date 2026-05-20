
/**
 * Samples for StorageAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountList.json
     */
    /**
     * Sample code: StorageAccountList.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountList(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().list(com.azure.core.util.Context.NONE);
    }
}
