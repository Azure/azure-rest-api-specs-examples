
/**
 * Samples for AmlFilesystems ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/
     * amlFilesystems_ListByResourceGroup.json
     */
    /**
     * Sample code: amlFilesystems_ListByResourceGroup.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void
        amlFilesystemsListByResourceGroup(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().listByResourceGroup("scgroup", com.azure.core.util.Context.NONE);
    }
}
