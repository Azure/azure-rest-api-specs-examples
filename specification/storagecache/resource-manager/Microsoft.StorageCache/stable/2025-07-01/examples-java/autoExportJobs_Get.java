
/**
 * Samples for AutoExportJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/autoExportJobs_Get.
     * json
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
