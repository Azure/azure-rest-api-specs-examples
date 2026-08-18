
/**
 * Samples for ExpressRouteCircuits GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExpressRouteMultiCloudCircuitGet.json
     */
    /**
     * Sample code: Get MultiCloud ExpressRouteCircuit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getMultiCloudExpressRouteCircuit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().getByResourceGroupWithResponse("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
