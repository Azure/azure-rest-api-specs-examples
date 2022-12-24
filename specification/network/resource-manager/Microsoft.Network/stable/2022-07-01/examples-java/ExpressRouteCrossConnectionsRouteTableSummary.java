import com.azure.core.util.Context;

/** Samples for ExpressRouteCrossConnections ListRoutesTableSummary. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRouteCrossConnectionsRouteTableSummary.json
     */
    /**
     * Sample code: GetExpressRouteCrossConnectionsRouteTableSummary.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCrossConnectionsRouteTableSummary(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCrossConnections()
            .listRoutesTableSummary(
                "CrossConnection-SiliconValley", "<circuitServiceKey>", "AzurePrivatePeering", "primary", Context.NONE);
    }
}
