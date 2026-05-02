
/**
 * Samples for AscOperations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/AscOperations_Get.json
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
