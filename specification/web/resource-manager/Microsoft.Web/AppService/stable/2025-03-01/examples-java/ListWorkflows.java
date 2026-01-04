
/**
 * Samples for WebApps ListInstanceWorkflowsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWorkflows.json
     */
    /**
     * Sample code: List the workflows Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheWorkflowsSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listInstanceWorkflowsSlot("testrg123", "testsite2",
            "staging", com.azure.core.util.Context.NONE);
    }
}
