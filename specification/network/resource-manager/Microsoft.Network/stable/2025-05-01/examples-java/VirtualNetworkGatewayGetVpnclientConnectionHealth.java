
/**
 * Samples for VirtualNetworkGateways GetVpnclientConnectionHealth.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayGetVpnclientConnectionHealth.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayVpnclientConnectionHealth.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getVirtualNetworkGatewayVpnclientConnectionHealth(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways()
            .getVpnclientConnectionHealth("p2s-vnet-test", "vpnp2sgw", com.azure.core.util.Context.NONE);
    }
}
