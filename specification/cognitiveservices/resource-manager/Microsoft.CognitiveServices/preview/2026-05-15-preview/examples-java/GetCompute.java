
/**
 * Samples for Computes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetCompute.json
     */
    /**
     * Sample code: GetCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().getWithResponse("rgcognitiveservices", "myAccount", "myCompute",
            com.azure.core.util.Context.NONE);
    }
}
