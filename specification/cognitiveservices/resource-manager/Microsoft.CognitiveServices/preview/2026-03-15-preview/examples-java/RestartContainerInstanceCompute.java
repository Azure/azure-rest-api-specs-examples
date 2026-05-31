
/**
 * Samples for Computes Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/RestartContainerInstanceCompute.json
     */
    /**
     * Sample code: RestartContainerInstanceCompute.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        restartContainerInstanceCompute(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.computes().restart("rgcognitiveservices", "myAccount", "myContainerInstance",
            com.azure.core.util.Context.NONE);
    }
}
