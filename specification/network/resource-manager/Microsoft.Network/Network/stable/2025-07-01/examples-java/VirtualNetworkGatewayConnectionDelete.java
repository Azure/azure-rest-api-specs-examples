
/**
 * Samples for VirtualNetworkGatewayConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionDelete.json
     */
    /**
     * Sample code: DeleteVirtualNetworkGatewayConnection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteVirtualNetworkGatewayConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().delete("rg1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
