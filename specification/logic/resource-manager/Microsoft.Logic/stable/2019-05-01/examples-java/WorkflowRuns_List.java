/** Samples for WorkflowRuns List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRuns_List.json
     */
    /**
     * Sample code: List workflow runs.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listWorkflowRuns(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRuns()
            .list("test-resource-group", "test-workflow", null, null, com.azure.core.util.Context.NONE);
    }
}
