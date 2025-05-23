
/**
 * Samples for RaiPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * GetRaiPolicy.json
     */
    /**
     * Sample code: GetRaiPolicy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getRaiPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().getWithResponse("resourceGroupName", "accountName", "raiPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
