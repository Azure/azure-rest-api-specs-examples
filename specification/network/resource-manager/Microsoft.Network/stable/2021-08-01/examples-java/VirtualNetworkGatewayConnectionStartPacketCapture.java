import com.azure.core.util.Context;

/** Samples for VirtualNetworkGatewayConnections StartPacketCapture. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewayConnectionStartPacketCapture.json
     */
    /**
     * Sample code: Start packet capture on virtual network gateway connection without filter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startPacketCaptureOnVirtualNetworkGatewayConnectionWithoutFilter(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGatewayConnections()
            .startPacketCapture("rg1", "vpngwcn1", null, Context.NONE);
    }
}
