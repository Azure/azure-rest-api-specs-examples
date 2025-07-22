
/**
 * Samples for RaiBlocklists List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ListBlocklists.json
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
