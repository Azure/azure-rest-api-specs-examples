
/**
 * Samples for ExpressRouteCircuitAuthorizations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitAuthorizationGet.json
     */
    /**
     * Sample code: Get ExpressRouteCircuit Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteCircuitAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitAuthorizations().getWithResponse("rg1", "circuitName",
            "authorizationName", com.azure.core.util.Context.NONE);
    }
}
