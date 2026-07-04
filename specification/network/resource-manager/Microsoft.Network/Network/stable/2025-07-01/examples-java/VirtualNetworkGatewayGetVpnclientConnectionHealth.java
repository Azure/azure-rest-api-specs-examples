
/**
 * Samples for VirtualNetworkGateways GetVpnclientConnectionHealth.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetVpnclientConnectionHealth.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayVpnclientConnectionHealth.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayVpnclientConnectionHealth(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getVpnclientConnectionHealth("p2s-vnet-test", "vpnp2sgw",
            com.azure.core.util.Context.NONE);
    }
}
