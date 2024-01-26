
/**
 * Samples for WorkflowTriggers ListCallbackUrl.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/WorkflowTriggers_ListCallbackUrl.json
     */
    /**
     * Sample code: Get the callback URL for a trigger.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheCallbackURLForATrigger(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowTriggers().listCallbackUrlWithResponse(
            "test-resource-group", "test-name", "test-workflow", "manual", com.azure.core.util.Context.NONE);
    }
}
