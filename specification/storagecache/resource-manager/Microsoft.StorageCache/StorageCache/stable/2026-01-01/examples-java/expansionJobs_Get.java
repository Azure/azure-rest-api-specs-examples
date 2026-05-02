
/**
 * Samples for ExpansionJobs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/expansionJobs_Get.json
     */
    /**
     * Sample code: expansionJobs_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void expansionJobsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.expansionJobs().getWithResponse("scgroup", "fs1", "expansionjob1", com.azure.core.util.Context.NONE);
    }
}
