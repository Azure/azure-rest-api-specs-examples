
/**
 * Samples for WorkflowTriggers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggers_Get.json
     */
    /**
     * Sample code: Get a workflow trigger.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAWorkflowTrigger(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowTriggers().getWithResponse("test-resource-group",
            "test-name", "test-workflow", "manual", com.azure.core.util.Context.NONE);
    }
}
