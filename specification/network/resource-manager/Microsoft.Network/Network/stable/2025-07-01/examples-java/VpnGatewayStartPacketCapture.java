
/**
 * Samples for VpnGateways StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on vpn gateway without filter.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        startPacketCaptureOnVpnGatewayWithoutFilter(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().startPacketCapture("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
