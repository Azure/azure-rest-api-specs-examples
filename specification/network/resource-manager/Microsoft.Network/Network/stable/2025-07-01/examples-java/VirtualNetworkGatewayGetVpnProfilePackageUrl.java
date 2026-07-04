
/**
 * Samples for VirtualNetworkGateways GetVpnProfilePackageUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetVpnProfilePackageUrl.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayVPNProfilePackageURL.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayVPNProfilePackageURL(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getVpnProfilePackageUrl("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
