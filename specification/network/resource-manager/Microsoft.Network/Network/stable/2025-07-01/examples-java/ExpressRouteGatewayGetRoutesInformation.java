
/**
 * Samples for ExpressRouteGateways GetRoutesInformation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteGatewayGetRoutesInformation.json
     */
    /**
     * Sample code: ExpressRouteGatewayGetRoutesInformation.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        expressRouteGatewayGetRoutesInformation(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways().getRoutesInformation("rg1", "ergw1", false,
            com.azure.core.util.Context.NONE);
    }
}
