
/**
 * Samples for RaiTopics List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListRaiTopics.json
     */
    /**
     * Sample code: ListRaiTopics.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listRaiTopics(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiTopics().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
