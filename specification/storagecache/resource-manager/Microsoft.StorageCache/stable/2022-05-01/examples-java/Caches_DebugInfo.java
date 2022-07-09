import com.azure.core.util.Context;

/** Samples for Caches DebugInfo. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/Caches_DebugInfo.json
     */
    /**
     * Sample code: Caches_DebugInfo.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesDebugInfo(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().debugInfo("scgroup", "sc", Context.NONE);
    }
}
