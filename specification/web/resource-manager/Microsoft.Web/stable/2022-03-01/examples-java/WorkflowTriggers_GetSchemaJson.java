import com.azure.core.util.Context;

/** Samples for WorkflowTriggers GetSchemaJson. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/WorkflowTriggers_GetSchemaJson.json
     */
    /**
     * Sample code: Get trigger schema.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTriggerSchema(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWorkflowTriggers()
            .getSchemaJsonWithResponse("testResourceGroup", "test-name", "testWorkflow", "testTrigger", Context.NONE);
    }
}
