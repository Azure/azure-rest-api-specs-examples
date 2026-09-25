
/**
 * Samples for ContextCacheContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheContainerCRUD/ContextCacheContainers_Delete.json
     */
    /**
     * Sample code: Delete a Context Cache Container.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteAContextCacheContainer(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCacheContainers().delete("testrg", "testaccount", "gpt4-prompts",
            com.azure.core.util.Context.NONE);
    }
}
