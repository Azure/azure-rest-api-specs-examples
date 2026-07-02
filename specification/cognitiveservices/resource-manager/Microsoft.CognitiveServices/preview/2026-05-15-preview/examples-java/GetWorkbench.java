
/**
 * Samples for Workbenches Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetWorkbench.json
     */
    /**
     * Sample code: GetWorkbench.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getWorkbench(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.workbenches().getWithResponse("rgcognitiveservices", "myAccount", "myProject", "myWorkbench",
            com.azure.core.util.Context.NONE);
    }
}
