
/**
 * Samples for AmlFilesystems ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/amlFilesystems_ListByResourceGroup.json
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
