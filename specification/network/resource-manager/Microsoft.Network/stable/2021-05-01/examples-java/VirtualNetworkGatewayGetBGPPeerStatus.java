import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways GetBgpPeerStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayGetBGPPeerStatus.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayBGPPeerStatus.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayBGPPeerStatus(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .getBgpPeerStatus("rg1", "vpngw", null, Context.NONE);
    }
}
