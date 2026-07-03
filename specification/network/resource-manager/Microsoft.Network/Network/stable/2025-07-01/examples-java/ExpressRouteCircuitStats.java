
/**
 * Samples for ExpressRouteCircuits GetStats.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitStats.json
     */
    /**
     * Sample code: Get ExpressRoute Circuit Traffic Stats.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteCircuitTrafficStats(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().getStatsWithResponse("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
