
/**
 * Samples for StorageTargets ListByCache.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/StorageTargets_ListByCache.json
     */
    /**
     * Sample code: StorageTargets_List.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().listByCache("scgroup", "sc1", com.azure.core.util.Context.NONE);
    }
}
