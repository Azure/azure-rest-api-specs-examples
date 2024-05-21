
/**
 * Samples for Caches Stop.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/Caches_Stop.json
     */
    /**
     * Sample code: Caches_Stop.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesStop(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().stop("scgroup", "sc", com.azure.core.util.Context.NONE);
    }
}
