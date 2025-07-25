
/**
 * Samples for VirtualNetworkGateways GetVpnclientIpsecParameters.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkGatewayGetVpnClientIpsecParameters.json
     */
    /**
     * Sample code: Get VirtualNetworkGateway VpnClientIpsecParameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getVirtualNetworkGatewayVpnClientIpsecParameters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().getVpnclientIpsecParameters("rg1",
            "vpngw", com.azure.core.util.Context.NONE);
    }
}
