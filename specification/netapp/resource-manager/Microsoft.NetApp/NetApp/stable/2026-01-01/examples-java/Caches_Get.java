
/**
 * Samples for Caches Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Caches_Get.json
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
