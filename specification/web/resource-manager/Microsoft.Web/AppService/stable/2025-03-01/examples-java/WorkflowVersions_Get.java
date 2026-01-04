
/**
 * Samples for WorkflowVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowVersions_Get.json
     */
    /**
     * Sample code: Get a workflow version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAWorkflowVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowVersions().getWithResponse("test-resource-group",
            "test-name", "test-workflow", "08586676824806722526", com.azure.core.util.Context.NONE);
    }
}
