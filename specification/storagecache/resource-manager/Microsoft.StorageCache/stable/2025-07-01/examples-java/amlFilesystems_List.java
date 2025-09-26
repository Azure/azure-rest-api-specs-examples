
/**
 * Samples for AmlFilesystems List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/amlFilesystems_List
     * .json
     */
    /**
     * Sample code: amlFilesystems_List.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void amlFilesystemsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().list(com.azure.core.util.Context.NONE);
    }
}
