import com.azure.core.util.Context;

/** Samples for Caches Flush. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-01-01/examples/Caches_Flush.json
     */
    /**
     * Sample code: Caches_Flush.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesFlush(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().flush("scgroup", "sc", Context.NONE);
    }
}
