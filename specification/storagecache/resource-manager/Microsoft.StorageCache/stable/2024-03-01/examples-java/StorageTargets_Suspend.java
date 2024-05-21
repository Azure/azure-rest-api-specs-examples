
/**
 * Samples for StorageTargetOperation Suspend.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/
     * StorageTargets_Suspend.json
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
