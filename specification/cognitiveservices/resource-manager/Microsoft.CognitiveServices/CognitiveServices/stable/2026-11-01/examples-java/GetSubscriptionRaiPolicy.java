
/**
 * Samples for SubscriptionRaiPolicy Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/GetSubscriptionRaiPolicy.json
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
