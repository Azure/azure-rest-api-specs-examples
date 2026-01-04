
/**
 * Samples for ExpressRouteCircuits ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * ExpressRouteCircuitListByResourceGroup.json
     */
    /**
     * Sample code: List ExpressRouteCircuits in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listExpressRouteCircuitsInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getExpressRouteCircuits().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
