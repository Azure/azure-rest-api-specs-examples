
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void operationsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
