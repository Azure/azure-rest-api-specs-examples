
/**
 * Samples for ExpressRouteCircuitAuthorizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitAuthorizationList.json
     */
    /**
     * Sample code: List ExpressRouteCircuit Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteCircuitAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitAuthorizations().list("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
