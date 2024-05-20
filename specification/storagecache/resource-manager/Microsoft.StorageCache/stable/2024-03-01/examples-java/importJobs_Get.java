
/**
 * Samples for ImportJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/importJobs_Get.json
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
