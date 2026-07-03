
/**
 * Samples for ExpressRouteCrossConnections ListRoutesTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionsRouteTable.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionsRouteTable.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getExpressRouteCrossConnectionsRouteTable(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().listRoutesTable("CrossConnection-SiliconValley",
            "<circuitServiceKey>", "AzurePrivatePeering", "primary", com.azure.core.util.Context.NONE);
    }
}
