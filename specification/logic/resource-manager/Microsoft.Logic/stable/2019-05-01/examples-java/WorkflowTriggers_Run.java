/** Samples for WorkflowTriggers Run. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_Run.json
     */
    /**
     * Sample code: Run a workflow trigger.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void runAWorkflowTrigger(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggers()
            .runWithResponse("test-resource-group", "test-workflow", "manual", com.azure.core.util.Context.NONE);
    }
}
