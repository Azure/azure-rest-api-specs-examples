
/**
 * Samples for ExpansionJobs ListByAmlFilesystem.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/expansionJobs_ListByAmlFilesystem.json
     */
    /**
     * Sample code: expansionJobs_ListByAmlFilesystem.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        expansionJobsListByAmlFilesystem(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.expansionJobs().listByAmlFilesystem("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
