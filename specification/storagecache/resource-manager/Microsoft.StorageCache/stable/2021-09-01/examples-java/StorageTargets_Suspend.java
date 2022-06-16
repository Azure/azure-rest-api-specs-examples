import com.azure.core.util.Context;

/** Samples for StorageTargetOperation Suspend. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_Suspend.json
     */
    /**
     * Sample code: StorageTargets_Suspend.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsSuspend(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargetOperations().suspend("scgroup", "sc", "st1", Context.NONE);
    }
}
