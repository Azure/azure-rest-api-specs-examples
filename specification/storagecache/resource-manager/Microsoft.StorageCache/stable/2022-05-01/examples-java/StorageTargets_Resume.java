import com.azure.core.util.Context;

/** Samples for StorageTargetOperation Resume. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2022-05-01/examples/StorageTargets_Resume.json
     */
    /**
     * Sample code: StorageTargets_Resume.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsResume(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().resume("scgroup", "sc", "st1", Context.NONE);
    }
}
