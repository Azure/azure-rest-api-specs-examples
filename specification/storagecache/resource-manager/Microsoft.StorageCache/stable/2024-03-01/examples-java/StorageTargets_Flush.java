
/**
 * Samples for StorageTargetOperation Flush.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/
     * StorageTargets_Flush.json
     */
    /**
     * Sample code: StorageTargets_Flush.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsFlush(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().flush("scgroup", "sc", "st1", com.azure.core.util.Context.NONE);
    }
}
