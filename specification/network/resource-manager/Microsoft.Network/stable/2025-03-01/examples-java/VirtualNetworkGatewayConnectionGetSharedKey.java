
/**
 * Samples for VirtualNetworkGatewayConnections GetSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayConnectionGetSharedKey.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayConnectionSharedKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getVirtualNetworkGatewayConnectionSharedKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayConnections().getSharedKeyWithResponse("rg1",
            "connS2S", com.azure.core.util.Context.NONE);
    }
}
