
/**
 * Samples for WorkspaceEvaluations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceEvaluations_Get_Succeeded.json
     */
    /**
     * Sample code: Read a Succeeded template evaluation independently of workspace scenarios.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void readASucceededTemplateEvaluationIndependentlyOfWorkspaceScenarios(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceEvaluations().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000003", com.azure.core.util.Context.NONE);
    }
}
