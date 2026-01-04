
/**
 * Samples for VirtualNetworkGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkGatewayDelete.
     * json
     */
    /**
     * Sample code: DeleteVirtualNetworkGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualNetworkGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().delete("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
