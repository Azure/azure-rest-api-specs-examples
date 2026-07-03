
/**
 * Samples for ExpressRouteCircuits List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitListBySubscription.json
     */
    /**
     * Sample code: List ExpressRouteCircuits in a subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listExpressRouteCircuitsInASubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().list(com.azure.core.util.Context.NONE);
    }
}
