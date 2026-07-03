
/**
 * Samples for VirtualNetworkGateways GetVpnclientIpsecParameters.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetVpnClientIpsecParameters.json
     */
    /**
     * Sample code: Get VirtualNetworkGateway VpnClientIpsecParameters.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayVpnClientIpsecParameters(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getVpnclientIpsecParameters("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
