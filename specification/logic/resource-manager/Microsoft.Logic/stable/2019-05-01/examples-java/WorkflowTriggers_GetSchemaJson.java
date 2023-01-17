/** Samples for WorkflowTriggers GetSchemaJson. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_GetSchemaJson.json
     */
    /**
     * Sample code: Get trigger schema.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getTriggerSchema(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowTriggers()
            .getSchemaJsonWithResponse(
                "testResourceGroup", "testWorkflow", "testTrigger", com.azure.core.util.Context.NONE);
    }
}
