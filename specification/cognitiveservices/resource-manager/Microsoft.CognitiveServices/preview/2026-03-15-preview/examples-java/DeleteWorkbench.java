
/**
 * Samples for Workbenches Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteWorkbench.json
     */
    /**
     * Sample code: DeleteWorkbench.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteWorkbench(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().delete("rgcognitiveservices", "myAccount", "myProject", "myWorkbench",
            com.azure.core.util.Context.NONE);
    }
}
