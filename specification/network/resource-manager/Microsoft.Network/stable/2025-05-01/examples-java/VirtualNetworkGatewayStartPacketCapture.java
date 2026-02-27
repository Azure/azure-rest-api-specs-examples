
/**
 * Samples for VirtualNetworkGateways StartPacketCapture.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on virtual network gateway without filter.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        startPacketCaptureOnVirtualNetworkGatewayWithoutFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().startPacketCapture("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
