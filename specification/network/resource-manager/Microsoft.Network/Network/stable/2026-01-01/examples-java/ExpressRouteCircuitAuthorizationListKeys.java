
/**
 * Samples for ExpressRouteCircuitAuthorizations ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteCircuitAuthorizationListKeys.json
     */
    /**
     * Sample code: List ExpressRouteCircuit Authorization Keys.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listExpressRouteCircuitAuthorizationKeys(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitAuthorizations().listKeysWithResponse("rg1", "circuitName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
