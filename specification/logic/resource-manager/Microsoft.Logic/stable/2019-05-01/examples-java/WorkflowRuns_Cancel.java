/** Samples for WorkflowRuns Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRuns_Cancel.json
     */
    /**
     * Sample code: Cancel a workflow run.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void cancelAWorkflowRun(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRuns()
            .cancelWithResponse(
                "test-resource-group",
                "test-workflow",
                "08586676746934337772206998657CU22",
                com.azure.core.util.Context.NONE);
    }
}
