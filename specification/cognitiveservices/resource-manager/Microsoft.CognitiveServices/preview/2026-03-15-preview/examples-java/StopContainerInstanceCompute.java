
/**
 * Samples for Computes Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/StopContainerInstanceCompute.json
     */
    /**
     * Sample code: StopContainerInstanceCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        stopContainerInstanceCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().stop("rgcognitiveservices", "myAccount", "myContainerInstance",
            com.azure.core.util.Context.NONE);
    }
}
