
/**
 * Samples for ExpressRouteCircuits GetCircuitLinkFailoverSingleTestDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitGetCircuitLinkFailoverSingleTestDetails.json
     */
    /**
     * Sample code: ExpressRouteCircuitGetCircuitLinkFailoverSingleTestDetails.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteCircuitGetCircuitLinkFailoverSingleTestDetails(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().getCircuitLinkFailoverSingleTestDetails("rg1", "circuit1",
            "Primary", "BgpDisconnect", "00000000-0000-0000-0000-000000000001", com.azure.core.util.Context.NONE);
    }
}
