
/**
 * Samples for ExpressRouteCrossConnections ListArpTable.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCrossConnectionsArpTable.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionsArpTable.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getExpressRouteCrossConnectionsArpTable(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCrossConnections().listArpTable("CrossConnection-SiliconValley",
            "<circuitServiceKey>", "AzurePrivatePeering", "primary", com.azure.core.util.Context.NONE);
    }
}
