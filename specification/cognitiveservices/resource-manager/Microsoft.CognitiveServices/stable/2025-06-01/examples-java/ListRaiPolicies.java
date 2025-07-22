
/**
 * Samples for RaiPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ListRaiPolicies.json
     */
    /**
     * Sample code: ListRaiPolicies.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listRaiPolicies(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
