
/**
 * Samples for WorkflowTriggers GetSchemaJson.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/WorkflowTriggers_GetSchemaJson.json
     */
    /**
     * Sample code: Get trigger schema.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getTriggerSchema(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWorkflowTriggers().getSchemaJsonWithResponse("testResourceGroup", "test-name",
            "testWorkflow", "testTrigger", com.azure.core.util.Context.NONE);
    }
}
