
/**
 * Samples for ContextCaches GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheCRUD/ContextCaches_Get.json
     */
    /**
     * Sample code: Get a Context Cache.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getAContextCache(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().getByResourceGroupWithResponse("testrg", "testaccount",
            com.azure.core.util.Context.NONE);
    }
}
