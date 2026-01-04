
/**
 * Samples for WebApps GetWorkflow.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetWorkflow.json
     */
    /**
     * Sample code: GET a workflow.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gETAWorkflow(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getWorkflowWithResponse("testrg123", "testsite2",
            "stateful1", com.azure.core.util.Context.NONE);
    }
}
