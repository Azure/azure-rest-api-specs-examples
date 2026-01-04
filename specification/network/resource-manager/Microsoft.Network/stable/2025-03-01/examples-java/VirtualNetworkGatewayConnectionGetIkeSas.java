
/**
 * Samples for VirtualNetworkGatewayConnections GetIkeSas.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayConnectionGetIkeSas.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayConnectionIkeSa.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayConnectionIkeSa(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayConnections().getIkeSas("rg1", "vpngwcn1",
            com.azure.core.util.Context.NONE);
    }
}
