/** Samples for WorkflowTriggers Reset. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_Reset.json
     */
    /**
     * Sample code: Reset trigger.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void resetTrigger(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggers()
            .resetWithResponse("testResourceGroup", "testWorkflow", "testTrigger", com.azure.core.util.Context.NONE);
    }
}
