
/**
 * Samples for StorageAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageAccountGetAsyncSkuConversionStatus.json
     */
    /**
     * Sample code: StorageAccountGetAsyncSkuConversionStatus.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        storageAccountGetAsyncSkuConversionStatus(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageAccounts().getByResourceGroupWithResponse("res9407", "sto8596", null,
            com.azure.core.util.Context.NONE);
    }
}
