
/**
 * Samples for Workbenches Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/StopWorkbench.json
     */
    /**
     * Sample code: StopWorkbench.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void stopWorkbench(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().stop("rgcognitiveservices", "myAccount", "myProject", "myWorkbench",
            com.azure.core.util.Context.NONE);
    }
}
