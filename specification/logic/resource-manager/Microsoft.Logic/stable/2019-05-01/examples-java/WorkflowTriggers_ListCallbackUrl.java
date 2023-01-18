/** Samples for WorkflowTriggers ListCallbackUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_ListCallbackUrl.json
     */
    /**
     * Sample code: Get the callback URL for a trigger.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getTheCallbackURLForATrigger(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggers()
            .listCallbackUrlWithResponse(
                "test-resource-group", "test-workflow", "manual", com.azure.core.util.Context.NONE);
    }
}
