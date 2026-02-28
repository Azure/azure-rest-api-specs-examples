
/**
 * Samples for ExpressRouteCircuits Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ExpressRouteCircuitDelete.
     * json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteExpressRouteCircuit(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuits().delete("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
