
/**
 * Samples for ExpressRouteCircuitConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitConnectionDelete.json
     */
    /**
     * Sample code: Delete ExpressRouteCircuit.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteExpressRouteCircuit(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitConnections().delete("rg1", "ExpressRouteARMCircuitA",
            "AzurePrivatePeering", "circuitConnectionUSAUS", com.azure.core.util.Context.NONE);
    }
}
