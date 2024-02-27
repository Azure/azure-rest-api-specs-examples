
/**
 * Samples for AmlFilesystems GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * amlFilesystems_Get.json
     */
    /**
     * Sample code: amlFilesystems_Get.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void amlFilesystemsGet(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().getByResourceGroupWithResponse("scgroup", "fs1", com.azure.core.util.Context.NONE);
    }
}
