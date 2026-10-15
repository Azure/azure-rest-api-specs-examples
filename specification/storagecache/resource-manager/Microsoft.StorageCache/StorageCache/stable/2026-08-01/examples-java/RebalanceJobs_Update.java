
import com.azure.resourcemanager.storagecache.models.RebalanceJobAdminStatus;
import com.azure.resourcemanager.storagecache.models.RebalanceJobUpdate;
import com.azure.resourcemanager.storagecache.models.RebalanceJobUpdateProperties;

/**
 * Samples for RebalanceJobs Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01/RebalanceJobs_Update.json
     */
    /**
     * Sample code: RebalanceJobs_Update.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void rebalanceJobsUpdate(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.rebalanceJobs().update("scgroup", "fs1", "expansionjob1-rebalance",
            new RebalanceJobUpdate()
                .withProperties(new RebalanceJobUpdateProperties().withAdminStatus(RebalanceJobAdminStatus.CANCEL)),
            com.azure.core.util.Context.NONE);
    }
}
