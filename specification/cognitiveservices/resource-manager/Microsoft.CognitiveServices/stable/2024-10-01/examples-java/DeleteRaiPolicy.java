
/**
 * Samples for RaiPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/
     * DeleteRaiPolicy.json
     */
    /**
     * Sample code: DeleteRaiPolicy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteRaiPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().delete("resourceGroupName", "accountName", "raiPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
