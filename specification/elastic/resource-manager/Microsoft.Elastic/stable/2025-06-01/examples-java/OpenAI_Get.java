
/**
 * Samples for OpenAI Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/OpenAI_Get.json
     */
    /**
     * Sample code: OpenAI_Get.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void openAIGet(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.openAIs().getWithResponse("myResourceGroup", "myMonitor", "default", com.azure.core.util.Context.NONE);
    }
}
