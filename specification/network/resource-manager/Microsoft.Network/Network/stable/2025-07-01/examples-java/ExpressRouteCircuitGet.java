
/**
 * Samples for ExpressRouteCircuits GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitGet.json
     */
    /**
     * Sample code: Get ExpressRouteCircuit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getExpressRouteCircuit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().getByResourceGroupWithResponse("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
