
/**
 * Samples for OpenAI Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/OpenAI_Delete.json
     */
    /**
     * Sample code: OpenAI_Delete.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void openAIDelete(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.openAIs().deleteWithResponse("myResourceGroup", "myMonitor", "default",
            com.azure.core.util.Context.NONE);
    }
}
