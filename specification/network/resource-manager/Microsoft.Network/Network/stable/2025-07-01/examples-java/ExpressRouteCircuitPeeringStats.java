
/**
 * Samples for ExpressRouteCircuits GetPeeringStats.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitPeeringStats.json
     */
    /**
     * Sample code: Get ExpressRoute Circuit Peering Traffic Stats.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getExpressRouteCircuitPeeringTrafficStats(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().getPeeringStatsWithResponse("rg1", "circuitName",
            "peeringName", com.azure.core.util.Context.NONE);
    }
}
