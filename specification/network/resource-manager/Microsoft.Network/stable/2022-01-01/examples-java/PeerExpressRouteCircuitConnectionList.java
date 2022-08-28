import com.azure.core.util.Context;

/** Samples for PeerExpressRouteCircuitConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/PeerExpressRouteCircuitConnectionList.json
     */
    /**
     * Sample code: List Peer ExpressRouteCircuit Connection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listPeerExpressRouteCircuitConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPeerExpressRouteCircuitConnections()
            .list("rg1", "ExpressRouteARMCircuitA", "AzurePrivatePeering", Context.NONE);
    }
}
