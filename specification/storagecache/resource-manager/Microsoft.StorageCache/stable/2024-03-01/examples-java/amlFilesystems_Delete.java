
/**
 * Samples for AmlFilesystems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/
     * amlFilesystems_Delete.json
     */
    /**
     * Sample code: amlFilesystems_Delete.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void amlFilesystemsDelete(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().delete("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
