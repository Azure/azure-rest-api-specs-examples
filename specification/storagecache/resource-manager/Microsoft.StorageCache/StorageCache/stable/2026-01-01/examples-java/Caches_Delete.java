
/**
 * Samples for Caches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_Delete.json
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
