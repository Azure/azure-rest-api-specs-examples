
import com.azure.resourcemanager.storage.models.ContextCacheContainerPropertiesUpdate;
import com.azure.resourcemanager.storage.models.ContextCacheContainerUpdate;

/**
 * Samples for ContextCacheContainers Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheContainerCRUD/ContextCacheContainers_Update.json
     */
    /**
     * Sample code: Update a Context Cache Container.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void updateAContextCacheContainer(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCacheContainers().update("testrg", "testaccount", "gpt4-prompts",
            new ContextCacheContainerUpdate().withProperties(new ContextCacheContainerPropertiesUpdate()
                .withDescription("Updated container for GPT-4 prompt caching").withTimeToLive(14)),
            com.azure.core.util.Context.NONE);
    }
}
