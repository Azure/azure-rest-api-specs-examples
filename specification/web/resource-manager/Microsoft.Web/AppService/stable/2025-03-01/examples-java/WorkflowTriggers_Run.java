
/**
 * Samples for WorkflowTriggers Run.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggers_Run.json
     */
    /**
     * Sample code: Run a workflow trigger.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void runAWorkflowTrigger(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowTriggers().run("test-resource-group", "test-name",
            "test-workflow", "recurrence", com.azure.core.util.Context.NONE);
    }
}
