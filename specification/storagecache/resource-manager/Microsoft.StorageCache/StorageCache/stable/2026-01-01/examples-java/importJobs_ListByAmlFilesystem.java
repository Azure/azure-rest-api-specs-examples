
/**
 * Samples for ImportJobs ListByAmlFilesystem.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/importJobs_ListByAmlFilesystem.json
     */
    /**
     * Sample code: importJobs_ListByAmlFilesystem.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        importJobsListByAmlFilesystem(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.importJobs().listByAmlFilesystem("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
