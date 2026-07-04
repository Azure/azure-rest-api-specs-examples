
/**
 * Samples for ExpressRouteCircuits StartCircuitLinkFailoverTest.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitStartCircuitLinkFailoverTest.json
     */
    /**
     * Sample code: ExpressRouteCircuitStartCircuitLinkFailoverTest.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteCircuitStartCircuitLinkFailoverTest(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().startCircuitLinkFailoverTest("rg1", "circuit1", "Primary",
            "BgpDisconnect", com.azure.core.util.Context.NONE);
    }
}
