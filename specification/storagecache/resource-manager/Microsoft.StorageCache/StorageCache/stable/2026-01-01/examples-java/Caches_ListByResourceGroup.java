
/**
 * Samples for Caches ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_ListByResourceGroup.json
     */
    /**
     * Sample code: Caches_ListByResourceGroup.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void cachesListByResourceGroup(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.caches().listByResourceGroup("scgroup", com.azure.core.util.Context.NONE);
    }
}
