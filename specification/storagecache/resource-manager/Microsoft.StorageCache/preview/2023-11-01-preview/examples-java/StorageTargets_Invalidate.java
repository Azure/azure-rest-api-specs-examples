
/**
 * Samples for StorageTargetOperation Invalidate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * StorageTargets_Invalidate.json
     */
    /**
     * Sample code: StorageTargets_Invalidate.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsInvalidate(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().invalidate("scgroup", "sc", "st1", com.azure.core.util.Context.NONE);
    }
}
