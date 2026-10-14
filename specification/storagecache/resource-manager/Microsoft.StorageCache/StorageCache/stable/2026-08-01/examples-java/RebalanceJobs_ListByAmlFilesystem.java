
/**
 * Samples for RebalanceJobs ListByAmlFilesystem.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/RebalanceJobs_ListByAmlFilesystem.json
     */
    /**
     * Sample code: RebalanceJobs_ListByAmlFilesystem.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        rebalanceJobsListByAmlFilesystem(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.rebalanceJobs().listByAmlFilesystem("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
