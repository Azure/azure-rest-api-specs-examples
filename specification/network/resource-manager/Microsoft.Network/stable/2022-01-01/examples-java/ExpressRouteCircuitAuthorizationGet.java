import com.azure.core.util.Context;

/** Samples for ExpressRouteCircuitAuthorizations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ExpressRouteCircuitAuthorizationGet.json
     */
    /**
     * Sample code: Get ExpressRouteCircuit Authorization.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getExpressRouteCircuitAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getExpressRouteCircuitAuthorizations()
            .getWithResponse("rg1", "circuitName", "authorizationName", Context.NONE);
    }
}
