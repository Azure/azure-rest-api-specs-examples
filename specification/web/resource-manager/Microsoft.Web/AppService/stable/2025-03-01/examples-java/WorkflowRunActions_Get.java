
/**
 * Samples for WorkflowRunActions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowRunActions_Get.
     * json
     */
    /**
     * Sample code: Get a workflow run action.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAWorkflowRunAction(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActions().getWithResponse("test-resource-group",
            "test-name", "test-workflow", "08586676746934337772206998657CU22", "HTTP",
            com.azure.core.util.Context.NONE);
    }
}
