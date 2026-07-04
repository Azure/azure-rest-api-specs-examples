
/**
 * Samples for ExpressRouteGateways GetFailoverSingleTestDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayGetFailoverSingleTestDetails.json
     */
    /**
     * Sample code: ExpressRouteGatewayGetFailoverSingleTestDetails.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteGatewayGetFailoverSingleTestDetails(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().getFailoverSingleTestDetails("rg1", "ergw1", "Vancouver",
            "00000000-0000-0000-0000-000000000001", com.azure.core.util.Context.NONE);
    }
}
