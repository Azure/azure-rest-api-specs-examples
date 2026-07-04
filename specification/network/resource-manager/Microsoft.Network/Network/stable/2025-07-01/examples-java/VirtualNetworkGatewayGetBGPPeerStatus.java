
/**
 * Samples for VirtualNetworkGateways GetBgpPeerStatus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetBGPPeerStatus.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayBGPPeerStatus.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkGatewayBGPPeerStatus(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getBgpPeerStatus("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
