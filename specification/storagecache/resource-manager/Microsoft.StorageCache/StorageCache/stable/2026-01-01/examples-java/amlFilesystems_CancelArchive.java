
/**
 * Samples for AmlFilesystems CancelArchive.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/amlFilesystems_CancelArchive.json
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
