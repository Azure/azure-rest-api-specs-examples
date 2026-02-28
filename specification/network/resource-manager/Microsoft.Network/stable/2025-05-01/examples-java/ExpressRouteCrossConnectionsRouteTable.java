
/**
 * Samples for ExpressRouteCrossConnections ListRoutesTable.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRouteCrossConnectionsRouteTable.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionsRouteTable.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCrossConnectionsRouteTable(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCrossConnections().listRoutesTable(
            "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", "primary",
            com.azure.core.util.Context.NONE);
    }
}
