
/**
 * Samples for AmlFilesystems CancelArchive.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/
     * amlFilesystems_CancelArchive.json
     */
    /**
     * Sample code: amlFilesystems_cancelArchive.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void amlFilesystemsCancelArchive(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.amlFilesystems().cancelArchiveWithResponse("scgroup", "sc", com.azure.core.util.Context.NONE);
    }
}
