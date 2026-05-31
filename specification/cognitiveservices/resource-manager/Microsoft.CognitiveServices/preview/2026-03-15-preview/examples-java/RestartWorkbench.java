
/**
 * Samples for Workbenches Restart.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/RestartWorkbench.json
     */
    /**
     * Sample code: RestartWorkbench.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void restartWorkbench(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().restart("rgcognitiveservices", "myAccount", "myProject", "myWorkbench",
            com.azure.core.util.Context.NONE);
    }
}
