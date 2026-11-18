
/**
 * Samples for RaiPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/GetRaiPolicyAcsWithAnnotationsAndEgress.json
     */
    /**
     * Sample code: Get an ACS policy with optional annotations and an egress companion.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getAnACSPolicyWithOptionalAnnotationsAndAnEgressCompanion(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().getWithResponse("resource-group", "safety-account", "annotated-egress-guard",
            com.azure.core.util.Context.NONE);
    }
}
