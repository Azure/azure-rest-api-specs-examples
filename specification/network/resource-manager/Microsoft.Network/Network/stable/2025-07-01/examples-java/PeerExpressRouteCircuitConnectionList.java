
/**
 * Samples for PeerExpressRouteCircuitConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PeerExpressRouteCircuitConnectionList.json
     */
    /**
     * Sample code: List Peer ExpressRouteCircuit Connection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPeerExpressRouteCircuitConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPeerExpressRouteCircuitConnections().list("rg1", "ExpressRouteARMCircuitA",
            "AzurePrivatePeering", com.azure.core.util.Context.NONE);
    }
}
