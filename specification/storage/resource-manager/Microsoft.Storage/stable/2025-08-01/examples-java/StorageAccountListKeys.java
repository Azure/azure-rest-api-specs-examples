
/**
 * Samples for StorageAccounts ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountListKeys.json
     */
    /**
     * Sample code: StorageAccountListKeys.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountListKeys(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().listKeysWithResponse("res418", "sto2220", null,
            com.azure.core.util.Context.NONE);
    }
}
