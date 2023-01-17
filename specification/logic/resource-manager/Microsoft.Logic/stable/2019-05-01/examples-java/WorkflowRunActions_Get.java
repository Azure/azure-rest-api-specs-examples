/** Samples for WorkflowRunActions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActions_Get.json
     */
    /**
     * Sample code: Get a workflow run action.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAWorkflowRunAction(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRunActions()
            .getWithResponse(
                "test-resource-group",
                "test-workflow",
                "08586676746934337772206998657CU22",
                "HTTP",
                com.azure.core.util.Context.NONE);
    }
}
