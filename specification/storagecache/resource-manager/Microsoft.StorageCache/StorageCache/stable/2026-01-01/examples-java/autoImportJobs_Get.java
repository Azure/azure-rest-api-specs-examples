
/**
 * Samples for AutoImportJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/autoImportJobs_Get.json
     */
    /**
     * Sample code: autoImportJobs_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void autoImportJobsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.autoImportJobs().getWithResponse("scgroup", "fs1", "autojob1", com.azure.core.util.Context.NONE);
    }
}
