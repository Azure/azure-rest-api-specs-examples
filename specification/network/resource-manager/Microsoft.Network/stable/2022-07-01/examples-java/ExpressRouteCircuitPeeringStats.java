import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuits GetPeeringStats. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/ExpressRouteCircuitPeeringStats.json
     */
    /**
     * Sample code: Get ExpressRoute Circuit Peering Traffic Stats.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCircuitPeeringTrafficStats(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuits()
            .getPeeringStatsWithResponse("rg1", "circuitName", "peeringName", Context.NONE);
    }
}
