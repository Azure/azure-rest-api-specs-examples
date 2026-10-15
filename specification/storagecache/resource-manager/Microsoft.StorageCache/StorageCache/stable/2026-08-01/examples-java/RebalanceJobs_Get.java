
/**
 * Samples for RebalanceJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/RebalanceJobs_Get.json
     */
    /**
     * Sample code: RebalanceJobs_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void rebalanceJobsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.rebalanceJobs().getWithResponse("scgroup", "fs1", "expansionjob1-rebalance",
            com.azure.core.util.Context.NONE);
    }
}
