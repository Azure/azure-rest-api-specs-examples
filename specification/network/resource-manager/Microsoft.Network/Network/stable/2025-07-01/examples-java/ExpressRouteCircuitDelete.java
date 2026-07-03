
/**
 * Samples for ExpressRouteCircuits Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteExpressRouteCircuit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().delete("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
