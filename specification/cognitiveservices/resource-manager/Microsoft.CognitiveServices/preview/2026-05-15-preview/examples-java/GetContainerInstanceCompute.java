
/**
 * Samples for Computes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetContainerInstanceCompute.json
     */
    /**
     * Sample code: GetContainerInstanceCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getContainerInstanceCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().getWithResponse("rgcognitiveservices", "myAccount", "myContainerInstance",
            com.azure.core.util.Context.NONE);
    }
}
