
/**
 * Samples for ExpressRouteCircuitConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * ExpressRouteCircuitConnectionDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteExpressRouteCircuit(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuitConnections().delete("rg1",
            "ExpressRouteARMCircuitA", "AzurePrivatePeering", "circuitConnectionUSAUS",
            com.azure.core.util.Context.NONE);
    }
}
