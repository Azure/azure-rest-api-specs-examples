
/**
 * Samples for FlowLogs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkWatcherFlowLogDelete.
     * json
     */
    /**
     * Sample code: Delete flow log.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFlowLog(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFlowLogs().delete("rg1", "nw1", "fl",
            com.azure.core.util.Context.NONE);
    }
}
