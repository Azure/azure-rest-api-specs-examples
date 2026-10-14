
/**
 * Samples for RebalanceJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/RebalanceJobs_Delete.json
     */
    /**
     * Sample code: RebalanceJobs_Delete.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void rebalanceJobsDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.rebalanceJobs().delete("scgroup", "fs1", "expansionjob1-rebalance", com.azure.core.util.Context.NONE);
    }
}
