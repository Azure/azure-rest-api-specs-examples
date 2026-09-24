
/**
 * Samples for ContextCacheContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheContainerCRUD/ContextCacheContainers_Get.json
     */
    /**
     * Sample code: Get a Context Cache Container.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void getAContextCacheContainer(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCacheContainers().getWithResponse("testrg", "testaccount", "gpt4-prompts",
            com.azure.core.util.Context.NONE);
    }
}
