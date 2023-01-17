/** Samples for WorkflowRuns Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRuns_Get.json
     */
    /**
     * Sample code: Get a run for a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getARunForAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowRuns()
            .getWithResponse(
                "test-resource-group",
                "test-workflow",
                "08586676746934337772206998657CU22",
                com.azure.core.util.Context.NONE);
    }
}
