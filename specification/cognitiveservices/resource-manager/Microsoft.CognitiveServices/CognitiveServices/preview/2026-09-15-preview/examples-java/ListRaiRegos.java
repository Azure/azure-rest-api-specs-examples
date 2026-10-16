
/**
 * Samples for RaiRegos List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/ListRaiRegos.json
     */
    /**
     * Sample code: List reusable Rego artifacts.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listReusableRegoArtifacts(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiRegos().list("resource-group", "safety-account", 10, com.azure.core.util.Context.NONE);
    }
}
