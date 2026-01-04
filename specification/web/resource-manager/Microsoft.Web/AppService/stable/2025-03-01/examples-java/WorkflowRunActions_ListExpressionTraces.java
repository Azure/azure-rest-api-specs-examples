
/**
 * Samples for WorkflowRunActions ListExpressionTraces.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * WorkflowRunActions_ListExpressionTraces.json
     */
    /**
     * Sample code: List expression traces.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressionTraces(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWorkflowRunActions().listExpressionTraces("testResourceGroup",
            "test-name", "testFlow", "08586776228332053161046300351", "testAction", com.azure.core.util.Context.NONE);
    }
}
