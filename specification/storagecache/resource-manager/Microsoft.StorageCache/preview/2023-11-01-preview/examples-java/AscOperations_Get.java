
/**
 * Samples for AscOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * AscOperations_Get.json
     */
    /**
     * Sample code: AscOperations_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void ascOperationsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.ascOperations().getWithResponse("westus", "testoperationid", com.azure.core.util.Context.NONE);
    }
}
