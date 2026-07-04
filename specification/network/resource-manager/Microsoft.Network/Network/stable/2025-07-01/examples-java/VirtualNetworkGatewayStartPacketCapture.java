
/**
 * Samples for VirtualNetworkGateways StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on virtual network gateway without filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void startPacketCaptureOnVirtualNetworkGatewayWithoutFilter(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().startPacketCapture("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
