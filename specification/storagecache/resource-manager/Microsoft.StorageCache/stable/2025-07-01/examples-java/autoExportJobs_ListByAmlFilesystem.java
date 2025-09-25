
/**
 * Samples for AutoExportJobs ListByAmlFilesystem.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/
     * autoExportJobs_ListByAmlFilesystem.json
     */
    /**
     * Sample code: autoExportJobs_ListByAmlFilesystem.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        autoExportJobsListByAmlFilesystem(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.autoExportJobs().listByAmlFilesystem("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
