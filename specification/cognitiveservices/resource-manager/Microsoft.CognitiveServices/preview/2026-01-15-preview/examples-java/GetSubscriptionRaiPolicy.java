
/**
 * Samples for SubscriptionRaiPolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetSubscriptionRaiPolicy.json
     */
    /**
     * Sample code: GetRaiPolicy.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getRaiPolicy(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.subscriptionRaiPolicies().getWithResponse("raiPolicyName", com.azure.core.util.Context.NONE);
    }
}
