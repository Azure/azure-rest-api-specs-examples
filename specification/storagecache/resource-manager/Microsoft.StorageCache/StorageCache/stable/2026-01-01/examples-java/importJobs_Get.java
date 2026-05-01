
/**
 * Samples for ImportJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/importJobs_Get.json
     */
    /**
     * Sample code: importJobs_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void importJobsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.importJobs().getWithResponse("scgroup", "fs1", "job1", com.azure.core.util.Context.NONE);
    }
}
