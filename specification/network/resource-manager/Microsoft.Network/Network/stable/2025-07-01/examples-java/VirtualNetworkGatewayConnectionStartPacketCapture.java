
/**
 * Samples for VirtualNetworkGatewayConnections StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayConnectionStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on virtual network gateway connection without filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void startPacketCaptureOnVirtualNetworkGatewayConnectionWithoutFilter(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayConnections().startPacketCapture("rg1", "vpngwcn1", null,
            com.azure.core.util.Context.NONE);
    }
}
