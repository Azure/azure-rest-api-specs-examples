
/**
 * Samples for StorageTargets Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * StorageTargets_Get.json
     */
    /**
     * Sample code: StorageTargets_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void storageTargetsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.storageTargets().getWithResponse("scgroup", "sc1", "st1", com.azure.core.util.Context.NONE);
    }
}
