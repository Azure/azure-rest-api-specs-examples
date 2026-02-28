
/**
 * Samples for ExpressRouteCircuitConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * ExpressRouteCircuitConnectionList.json
     */
    /**
     * Sample code: List ExpressRouteCircuit Connection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressRouteCircuitConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuitConnections().list("rg1",
            "ExpressRouteARMCircuitA", "AzurePrivatePeering", com.azure.core.util.Context.NONE);
    }
}
