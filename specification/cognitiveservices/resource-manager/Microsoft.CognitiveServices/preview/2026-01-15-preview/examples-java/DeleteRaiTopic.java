
/**
 * Samples for RaiTopics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/DeleteRaiTopic.json
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
