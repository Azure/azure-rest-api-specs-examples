
/**
 * Samples for AutoImportJobs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/autoImportJobs_Delete.json
     */
    /**
     * Sample code: autoImportJobs_Delete.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void autoImportJobsDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.autoImportJobs().delete("scgroup", "fs1", "autojob1", com.azure.core.util.Context.NONE);
    }
}
