
/**
 * Samples for AutoExportJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/autoExportJobs_Get.json
     */
    /**
     * Sample code: autoExportJobs_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void autoExportJobsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.autoExportJobs().getWithResponse("scgroup", "fs1", "job1", com.azure.core.util.Context.NONE);
    }
}
