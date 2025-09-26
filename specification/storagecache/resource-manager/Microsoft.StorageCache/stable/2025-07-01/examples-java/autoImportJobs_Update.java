
import com.azure.resourcemanager.storagecache.models.AutoImportJob;
import com.azure.resourcemanager.storagecache.models.AutoImportJobUpdatePropertiesAdminStatus;

/**
 * Samples for AutoImportJobs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/
     * autoImportJobs_Update.json
     */
    /**
     * Sample code: autoImportJobs_Update.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void autoImportJobsUpdate(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        AutoImportJob resource = manager.autoImportJobs()
            .getWithResponse("scgroup", "fs1", "autojob1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withAdminStatus(AutoImportJobUpdatePropertiesAdminStatus.DISABLE).apply();
    }
}
