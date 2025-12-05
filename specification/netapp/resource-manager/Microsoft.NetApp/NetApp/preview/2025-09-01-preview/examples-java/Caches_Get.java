
/**
 * Samples for Caches Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Caches_Get.json
     */
    /**
     * Sample code: Caches_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().getWithResponse("myRG", "account1", "pool1", "cache1", com.azure.core.util.Context.NONE);
    }
}
