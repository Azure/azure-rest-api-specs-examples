
/**
 * Samples for ExpressRouteCircuits GetCircuitLinkFailoverAllTestsDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitGetCircuitLinkFailoverAllTestsDetails.json
     */
    /**
     * Sample code: ExpressRouteCircuitGetCircuitLinkFailoverAllTestsDetails.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteCircuitGetCircuitLinkFailoverAllTestsDetails(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().getCircuitLinkFailoverAllTestsDetails("rg1", "circuit1",
            "SingleSiteFailover", true, com.azure.core.util.Context.NONE);
    }
}
