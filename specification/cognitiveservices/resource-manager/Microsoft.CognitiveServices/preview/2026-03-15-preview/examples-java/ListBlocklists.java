
/**
 * Samples for RaiBlocklists List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListBlocklists.json
     */
    /**
     * Sample code: ListBlocklists.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listBlocklists(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklists().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
