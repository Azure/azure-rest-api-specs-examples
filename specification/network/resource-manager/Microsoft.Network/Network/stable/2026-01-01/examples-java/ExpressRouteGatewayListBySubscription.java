
/**
 * Samples for ExpressRouteGateways ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteGatewayListBySubscription.json
     */
    /**
     * Sample code: ExpressRouteGatewayListBySubscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void expressRouteGatewayListBySubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteGateways()
            .listBySubscriptionWithResponse(com.azure.core.util.Context.NONE);
    }
}
