
/**
 * Samples for OpenAI CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/OpenAI_CreateOrUpdate.json
     */
    /**
     * Sample code: OpenAI_CreateOrUpdate.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void openAICreateOrUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.openAIs().define("default").withExistingMonitor("myResourceGroup", "myMonitor").create();
    }
}
