
/**
 * Samples for WorkspaceEvaluations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/WorkspaceEvaluations_Get_PartiallySucceeded.json
     */
    /**
     * Sample code: Read a PartiallySucceeded template evaluation independently of workspace scenarios.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void readAPartiallySucceededTemplateEvaluationIndependentlyOfWorkspaceScenarios(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.workspaceEvaluations().getWithResponse("exampleRG", "exampleWorkspace",
            "00000000-0000-0000-0000-000000000003", com.azure.core.util.Context.NONE);
    }
}
