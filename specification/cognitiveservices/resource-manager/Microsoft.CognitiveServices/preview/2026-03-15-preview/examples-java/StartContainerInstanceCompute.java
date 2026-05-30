
/**
 * Samples for Computes Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/StartContainerInstanceCompute.json
     */
    /**
     * Sample code: StartContainerInstanceCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        startContainerInstanceCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().start("rgcognitiveservices", "myAccount", "myContainerInstance",
            com.azure.core.util.Context.NONE);
    }
}
