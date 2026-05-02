
/**
 * Samples for AutoExportJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/autoExportJobs_Delete.json
     */
    /**
     * Sample code: autoExportJobs_Delete.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void autoExportJobsDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.autoExportJobs().delete("scgroup", "fs1", "job1", com.azure.core.util.Context.NONE);
    }
}
