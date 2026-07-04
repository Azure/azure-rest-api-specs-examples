
/**
 * Samples for ExpressRouteGateways StartSiteFailoverTest.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayStartSiteFailoverTest.json
     */
    /**
     * Sample code: ExpressRouteGatewayStartSiteFailoverTest.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteGatewayStartSiteFailoverTest(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().startSiteFailoverTest("rg1", "ergw1", "Vancouver",
            com.azure.core.util.Context.NONE);
    }
}
