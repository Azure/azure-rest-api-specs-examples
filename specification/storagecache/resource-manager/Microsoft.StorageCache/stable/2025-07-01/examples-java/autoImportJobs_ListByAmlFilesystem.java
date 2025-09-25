
/**
 * Samples for AutoImportJobs ListByAmlFilesystem.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/
     * autoImportJobs_ListByAmlFilesystem.json
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
