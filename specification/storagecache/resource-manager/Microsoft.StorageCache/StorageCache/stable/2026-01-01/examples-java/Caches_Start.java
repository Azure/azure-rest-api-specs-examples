
/**
 * Samples for Caches Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_Start.json
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
