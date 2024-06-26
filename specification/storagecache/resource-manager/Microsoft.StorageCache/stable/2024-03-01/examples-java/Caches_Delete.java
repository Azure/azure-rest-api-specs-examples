
/**
 * Samples for Caches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/Caches_Delete.json
     */
    /**
     * Sample code: Caches_Delete.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().delete("scgroup", "sc", com.azure.core.util.Context.NONE);
    }
}
