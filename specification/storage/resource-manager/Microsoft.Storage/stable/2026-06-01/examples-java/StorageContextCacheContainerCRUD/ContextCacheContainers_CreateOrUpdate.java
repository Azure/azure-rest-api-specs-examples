
import com.azure.resourcemanager.storage.fluent.models.ContextCacheContainerInner;
import com.azure.resourcemanager.storage.models.AiProvider;
import com.azure.resourcemanager.storage.models.ContextCacheContainerProperties;

/**
 * Samples for ContextCacheContainers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/StorageContextCacheContainerCRUD/ContextCacheContainers_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a Context Cache Container.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createAContextCacheContainer(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getContextCacheContainers().createOrUpdate("testrg", "testaccount", "gpt4-prompts",
            new ContextCacheContainerInner().withProperties(
                new ContextCacheContainerProperties().withDescription("Container for GPT-4 prompt caching")
                    .withModelName("gpt-4").withProvider(AiProvider.OPEN_AI).withTimeToLive(7)),
            com.azure.core.util.Context.NONE);
    }
}
