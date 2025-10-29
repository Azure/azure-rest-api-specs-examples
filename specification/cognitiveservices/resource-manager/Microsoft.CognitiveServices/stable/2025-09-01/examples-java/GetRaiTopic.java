
/**
 * Samples for RaiTopics Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * GetRaiTopic.json
     */
    /**
     * Sample code: GetRaiTopic.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getRaiTopic(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiTopics().getWithResponse("resourceGroupName", "accountName", "raiTopicName",
            com.azure.core.util.Context.NONE);
    }
}
