
/**
 * Samples for VirtualNetworkGatewayConnections GetSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionGetSharedKey.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayConnectionSharedKey.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayConnectionSharedKey(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().getSharedKeyWithResponse("rg1", "connS2S",
            com.azure.core.util.Context.NONE);
    }
}
