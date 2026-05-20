
/**
 * Samples for StorageAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/StorageAccountGetProperties.json
     */
    /**
     * Sample code: StorageAccountGetProperties.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountGetProperties(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().getByResourceGroupWithResponse("res9407", "sto8596", null,
            com.azure.core.util.Context.NONE);
    }
}
