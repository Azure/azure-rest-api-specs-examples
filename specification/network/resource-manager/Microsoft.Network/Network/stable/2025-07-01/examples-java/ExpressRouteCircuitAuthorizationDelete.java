
/**
 * Samples for ExpressRouteCircuitAuthorizations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitAuthorizationDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteExpressRouteCircuitAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitAuthorizations().delete("rg1", "circuitName", "authorizationName",
            com.azure.core.util.Context.NONE);
    }
}
