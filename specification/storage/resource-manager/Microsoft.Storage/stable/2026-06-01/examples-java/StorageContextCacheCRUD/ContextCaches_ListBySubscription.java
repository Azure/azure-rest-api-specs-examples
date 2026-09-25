
/**
 * Samples for ContextCaches List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheCRUD/ContextCaches_ListBySubscription.json
     */
    /**
     * Sample code: List Context Caches by Subscription.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void listContextCachesBySubscription(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().list(com.azure.core.util.Context.NONE);
    }
}
