
/**
 * Samples for RaiRegos Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/GetRaiRego.json
     */
    /**
     * Sample code: Get a reusable Rego artifact.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAReusableRegoArtifact(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiRegos().getWithResponse("resource-group", "safety-account", "input-guard",
            com.azure.core.util.Context.NONE);
    }
}
