
/**
 * Samples for Caches List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/Caches_List.json
     */
    /**
     * Sample code: Caches_List.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().list(com.azure.core.util.Context.NONE);
    }
}
