/** Samples for StorageTargetOperation Resume. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-03-01-preview/examples/StorageTargets_Resume.json
     */
    /**
     * Sample code: StorageTargets_Resume.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsResume(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().resume("scgroup", "sc", "st1", com.azure.core.util.Context.NONE);
    }
}
