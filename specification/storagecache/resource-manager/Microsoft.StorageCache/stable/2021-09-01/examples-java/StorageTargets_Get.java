import com.azure.core.util.Context;

/** Samples for StorageTargets Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2021-09-01/examples/StorageTargets_Get.json
     */
    /**
     * Sample code: StorageTargets_Get.
     *
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().getWithResponse("scgroup", "sc1", "st1", Context.NONE);
    }
}
