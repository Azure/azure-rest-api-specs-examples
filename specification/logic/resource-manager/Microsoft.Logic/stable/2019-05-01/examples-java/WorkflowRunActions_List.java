/** Samples for WorkflowRunActions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActions_List.json
     */
    /**
     * Sample code: List a workflow run actions.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listAWorkflowRunActions(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRunActions()
            .list(
                "test-resource-group",
                "test-workflow",
                "08586676746934337772206998657CU22",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
