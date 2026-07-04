
/**
 * Samples for VirtualNetworkGatewayConnections ResetConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionReset.json
     */
    /**
     * Sample code: ResetVirtualNetworkGatewayConnection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void resetVirtualNetworkGatewayConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().resetConnection("rg1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
