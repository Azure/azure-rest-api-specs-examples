
/**
 * Samples for RaiTopics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * DeleteRaiTopic.json
     */
    /**
     * Sample code: DeleteRaiTopic.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteRaiTopic(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiTopics().delete("resourceGroupName", "accountName", "raiTopicName",
            com.azure.core.util.Context.NONE);
    }
}
