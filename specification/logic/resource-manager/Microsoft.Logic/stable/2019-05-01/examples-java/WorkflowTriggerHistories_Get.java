/** Samples for WorkflowTriggerHistories Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_Get.json
     */
    /**
     * Sample code: Get a workflow trigger history.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAWorkflowTriggerHistory(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggerHistories()
            .getWithResponse(
                "testResourceGroup",
                "testWorkflowName",
                "testTriggerName",
                "08586676746934337772206998657CU22",
                com.azure.core.util.Context.NONE);
    }
}
