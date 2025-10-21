
/**
 * Samples for OpenAI List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2025-06-01/examples/OpenAI_List.json
     */
    /**
     * Sample code: OpenAI_List.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void openAIList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.openAIs().list("myResourceGroup", "myMonitor", com.azure.core.util.Context.NONE);
    }
}
