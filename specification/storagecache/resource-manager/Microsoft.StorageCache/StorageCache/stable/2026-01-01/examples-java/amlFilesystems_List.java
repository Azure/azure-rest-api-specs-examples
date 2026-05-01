
/**
 * Samples for AmlFilesystems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/amlFilesystems_List.json
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
