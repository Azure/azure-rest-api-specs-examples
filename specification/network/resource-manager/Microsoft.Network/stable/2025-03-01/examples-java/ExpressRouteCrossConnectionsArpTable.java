
/**
 * Samples for ExpressRouteCrossConnections ListArpTable.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRouteCrossConnectionsArpTable.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionsArpTable.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCrossConnectionsArpTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnections().listArpTable(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", "primary",
            com.azure.core.util.Context.NONE);
    }
}
