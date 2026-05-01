
/**
 * Samples for UsageModels List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/UsageModels_List.json
     */
    /**
     * Sample code: UsageModels_List.
     * 
     * @param manager Entry point to StorageCacheManager.
     */
    public static void usageModelsList(com.azure.resourcemanager.storagecache.StorageCacheManager manager) {
        manager.usageModels().list(com.azure.core.util.Context.NONE);
    }
}
