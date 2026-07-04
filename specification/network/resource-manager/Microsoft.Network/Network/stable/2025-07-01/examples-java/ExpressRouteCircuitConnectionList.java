
/**
 * Samples for ExpressRouteCircuitConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ExpressRouteCircuitConnectionList.json
     */
    /**
     * Sample code: List ExpressRouteCircuit Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listExpressRouteCircuitConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteCircuitConnections().list("rg1", "ExpressRouteARMCircuitA",
            "AzurePrivatePeering", com.azure.core.util.Context.NONE);
    }
}
