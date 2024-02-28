
/**
 * Samples for Caches Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * Caches_Start.json
     */
    /**
     * Sample code: Caches_Start.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesStart(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().start("scgroup", "sc", com.azure.core.util.Context.NONE);
    }
}
