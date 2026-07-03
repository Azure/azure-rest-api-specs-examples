
/**
 * Samples for ExpressRouteCircuits ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitListByResourceGroup.json
     */
    /**
     * Sample code: List ExpressRouteCircuits in a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listExpressRouteCircuitsInAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
