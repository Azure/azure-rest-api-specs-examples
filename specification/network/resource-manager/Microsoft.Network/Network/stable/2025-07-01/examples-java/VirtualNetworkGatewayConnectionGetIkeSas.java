
/**
 * Samples for VirtualNetworkGatewayConnections GetIkeSas.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionGetIkeSas.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayConnectionIkeSa.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayConnectionIkeSa(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().getIkeSas("rg1", "vpngwcn1",
            com.azure.core.util.Context.NONE);
    }
}
