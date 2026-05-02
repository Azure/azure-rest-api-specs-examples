
/**
 * Samples for AutoImportJobs ListByAmlFilesystem.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/autoImportJobs_ListByAmlFilesystem.json
     */
    /**
     * Sample code: autoImportJobs_ListByAmlFilesystem.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        autoImportJobsListByAmlFilesystem(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.autoImportJobs().listByAmlFilesystem("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
