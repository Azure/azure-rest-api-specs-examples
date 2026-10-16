
/**
 * Samples for RaiPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/GetRaiPolicyAcs.json
     */
    /**
     * Sample code: Get an ACS policy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getAnACSPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().getWithResponse("resource-group", "safety-account", "agent-guard",
            com.azure.core.util.Context.NONE);
    }
}
