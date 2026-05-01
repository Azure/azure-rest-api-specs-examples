
/**
 * Samples for Caches DebugInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_DebugInfo.json
     */
    /**
     * Sample code: Caches_DebugInfo.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesDebugInfo(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().debugInfo("scgroup", "sc", com.azure.core.util.Context.NONE);
    }
}
