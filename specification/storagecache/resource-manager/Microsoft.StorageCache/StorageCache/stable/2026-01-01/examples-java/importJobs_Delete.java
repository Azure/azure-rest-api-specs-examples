
/**
 * Samples for ImportJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/importJobs_Delete.json
     */
    /**
     * Sample code: importJobs_Delete.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void importJobsDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.importJobs().delete("scgroup", "fs1", "job1", com.azure.core.util.Context.NONE);
    }
}
