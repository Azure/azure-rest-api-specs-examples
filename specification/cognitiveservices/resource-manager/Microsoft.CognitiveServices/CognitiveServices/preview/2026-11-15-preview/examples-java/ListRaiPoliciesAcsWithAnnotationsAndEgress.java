
/**
 * Samples for RaiPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/ListRaiPoliciesAcsWithAnnotationsAndEgress.json
     */
    /**
     * Sample code: List ACS policies with optional annotations and egress companions.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listACSPoliciesWithOptionalAnnotationsAndEgressCompanions(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiPolicies().list("resource-group", "safety-account", com.azure.core.util.Context.NONE);
    }
}
