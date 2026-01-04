
/**
 * Samples for VirtualNetworkGatewayConnections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayConnectionsList.json
     */
    /**
     * Sample code: ListVirtualNetworkGatewayConnectionsinResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listVirtualNetworkGatewayConnectionsinResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayConnections().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
