
/**
 * Samples for RaiPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/DeleteRaiPolicyAcs.json
     */
    /**
     * Sample code: Delete an ACS policy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteAnACSPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().delete("resource-group", "safety-account", "agent-guard",
            com.azure.core.util.Context.NONE);
    }
}
