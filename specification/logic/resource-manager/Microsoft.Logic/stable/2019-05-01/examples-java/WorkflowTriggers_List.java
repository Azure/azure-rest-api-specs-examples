/** Samples for WorkflowTriggers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_List.json
     */
    /**
     * Sample code: List workflow triggers.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listWorkflowTriggers(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggers()
            .list("test-resource-group", "test-workflow", null, null, com.azure.core.util.Context.NONE);
    }
}
