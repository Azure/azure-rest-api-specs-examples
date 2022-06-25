import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuits GetStats. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ExpressRouteCircuitStats.json
     */
    /**
     * Sample code: Get ExpressRoute Circuit Traffic Stats.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCircuitTrafficStats(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuits()
            .getStatsWithResponse("rg1", "circuitName", Context.NONE);
    }
}
