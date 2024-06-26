
/**
 * Samples for Caches Flush.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * Caches_Flush.json
     */
    /**
     * Sample code: Caches_Flush.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesFlush(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().flush("scgroup", "sc", com.azure.core.util.Context.NONE);
    }
}
