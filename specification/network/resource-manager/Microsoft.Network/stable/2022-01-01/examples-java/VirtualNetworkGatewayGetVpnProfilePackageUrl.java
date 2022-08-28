import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways GetVpnProfilePackageUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayGetVpnProfilePackageUrl.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayVPNProfilePackageURL.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayVPNProfilePackageURL(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .getVpnProfilePackageUrl("rg1", "vpngw", Context.NONE);
    }
}
