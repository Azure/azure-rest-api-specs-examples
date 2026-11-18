
/**
 * Samples for ProjectCostControls Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/ProjectCostControl/get.json
     */
    /**
     * Sample code: Get a project-owned budget.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAProjectOwnedBudget(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCostControls().getWithResponse("foundry-resource-group", "foundry-account", "production",
            "project-token-budget", com.azure.core.util.Context.NONE);
    }
}
