
/**
 * Samples for ContextCaches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheCRUD/ContextCaches_Delete.json
     */
    /**
     * Sample code: Delete a Context Cache.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteAContextCache(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCaches().delete("testrg", "testaccount", com.azure.core.util.Context.NONE);
    }
}
