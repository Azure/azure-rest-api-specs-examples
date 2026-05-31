
/**
 * Samples for SubscriptionRaiPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteSubscriptionRaiPolicy.json
     */
    /**
     * Sample code: DeleteRaiPolicy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteRaiPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.subscriptionRaiPolicies().delete("raiPolicyName", com.azure.core.util.Context.NONE);
    }
}
