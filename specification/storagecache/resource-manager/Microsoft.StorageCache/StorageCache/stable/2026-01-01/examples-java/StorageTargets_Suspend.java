
/**
 * Samples for StorageTargetOperation Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/StorageTargets_Suspend.json
     */
    /**
     * Sample code: StorageTargets_Suspend.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsSuspend(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().suspend("scgroup", "sc", "st1", com.azure.core.util.Context.NONE);
    }
}
