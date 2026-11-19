
/**
 * Samples for ProjectCostControls List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-15-preview/ProjectCostControl/list.json
     */
    /**
     * Sample code: List budgets owned by a project.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listBudgetsOwnedByAProject(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.projectCostControls().list("foundry-resource-group", "foundry-account", "production",
            com.azure.core.util.Context.NONE);
    }
}
