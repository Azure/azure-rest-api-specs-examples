
/**
 * Samples for FlowLogs Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkWatcherFlowLogGet.json
     */
    /**
     * Sample code: Get flow log.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFlowLog(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFlowLogs().getWithResponse("rg1", "nw1", "flowLog1",
            com.azure.core.util.Context.NONE);
    }
}
