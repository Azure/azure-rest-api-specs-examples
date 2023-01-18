/** Samples for WorkflowTriggerHistories List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_List.json
     */
    /**
     * Sample code: List a workflow trigger history.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listAWorkflowTriggerHistory(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggerHistories()
            .list(
                "testResourceGroup",
                "testWorkflowName",
                "testTriggerName",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
