
/**
 * Samples for Workbenches Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/StartWorkbench.json
     */
    /**
     * Sample code: StartWorkbench.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void startWorkbench(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().start("rgcognitiveservices", "myAccount", "myProject", "myWorkbench",
            com.azure.core.util.Context.NONE);
    }
}
