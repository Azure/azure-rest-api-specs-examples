
/**
 * Samples for AmlFilesystems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/amlFilesystems_Delete.json
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
