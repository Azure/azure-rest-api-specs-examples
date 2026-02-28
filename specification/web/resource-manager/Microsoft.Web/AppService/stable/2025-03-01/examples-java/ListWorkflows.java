
/**
 * Samples for WebApps ListWorkflows.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWorkflows.json
     */
    /**
     * Sample code: List the workflows.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheWorkflows(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listWorkflows("testrg123", "testsite2",
            com.azure.core.util.Context.NONE);
    }
}
