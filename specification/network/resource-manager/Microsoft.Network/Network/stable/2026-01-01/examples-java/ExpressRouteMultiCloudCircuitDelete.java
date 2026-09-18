
/**
 * Samples for ExpressRouteCircuits Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ExpressRouteMultiCloudCircuitDelete.json
     */
    /**
     * Sample code: Delete MultiCloud ExpressRouteCircuit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteMultiCloudExpressRouteCircuit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuits().delete("rg1", "circuitName",
            com.azure.core.util.Context.NONE);
    }
}
