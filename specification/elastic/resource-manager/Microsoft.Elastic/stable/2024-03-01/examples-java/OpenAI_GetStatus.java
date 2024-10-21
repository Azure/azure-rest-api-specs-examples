
/**
 * Samples for OpenAI GetStatus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/OpenAI_GetStatus.json
     */
    /**
     * Sample code: OpenAI_GetStatus.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void openAIGetStatus(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.openAIs().getStatusWithResponse("myResourceGroup", "myMonitor", "default",
            com.azure.core.util.Context.NONE);
    }
}
