
/**
 * Samples for ExpressRouteCircuitAuthorizations List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRouteCircuitAuthorizationList.json
     */
    /**
     * Sample code: List ExpressRouteCircuit Authorization.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressRouteCircuitAuthorization(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuitAuthorizations().list("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
