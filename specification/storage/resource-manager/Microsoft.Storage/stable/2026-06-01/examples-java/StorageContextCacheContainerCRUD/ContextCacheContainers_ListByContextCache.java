
/**
 * Samples for ContextCacheContainers ListByContextCache.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheContainerCRUD/ContextCacheContainers_ListByContextCache.json
     */
    /**
     * Sample code: List Context Cache Containers in a Context Cache.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        listContextCacheContainersInAContextCache(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCacheContainers().listByContextCache("testrg", "testaccount",
            com.azure.core.util.Context.NONE);
    }
}
