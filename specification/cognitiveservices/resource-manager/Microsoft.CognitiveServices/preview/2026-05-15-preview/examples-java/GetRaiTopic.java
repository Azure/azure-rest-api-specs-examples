
/**
 * Samples for RaiTopics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetRaiTopic.json
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
