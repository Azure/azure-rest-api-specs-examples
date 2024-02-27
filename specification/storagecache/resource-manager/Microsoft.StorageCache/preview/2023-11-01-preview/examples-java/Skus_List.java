
/**
 * Samples for Skus List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/Skus_List.
     * json
     */
    /**
     * Sample code: Skus_List.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void skusList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.skus().list(com.azure.core.util.Context.NONE);
    }
}
