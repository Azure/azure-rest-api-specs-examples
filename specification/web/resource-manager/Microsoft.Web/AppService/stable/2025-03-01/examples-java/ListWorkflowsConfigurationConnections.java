
/**
 * Samples for WebApps ListWorkflowsConnectionsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListWorkflowsConfigurationConnections.json
     */
    /**
     * Sample code: List the Instance Workflows Configuration Connections Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listTheInstanceWorkflowsConfigurationConnectionsSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listWorkflowsConnectionsSlotWithResponse("testrg123",
            "testsite2", "staging", com.azure.core.util.Context.NONE);
    }
}
