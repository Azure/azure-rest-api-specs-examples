
/**
 * Samples for VpnGateways StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VpnGatewayStartPacketCapture.
     * json
     */
    /**
     * Sample code: Start packet capture on vpn gateway without filter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        startPacketCaptureOnVpnGatewayWithoutFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnGateways().startPacketCapture("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
