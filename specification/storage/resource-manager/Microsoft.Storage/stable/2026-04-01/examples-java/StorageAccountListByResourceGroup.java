
/**
 * Samples for StorageAccounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountListByResourceGroup.json
     */
    /**
     * Sample code: StorageAccountListByResourceGroup.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountListByResourceGroup(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().listByResourceGroup("res6117", com.azure.core.util.Context.NONE);
    }
}
