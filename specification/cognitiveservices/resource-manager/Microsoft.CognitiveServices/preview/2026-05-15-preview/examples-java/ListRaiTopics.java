
/**
 * Samples for RaiTopics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListRaiTopics.json
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
