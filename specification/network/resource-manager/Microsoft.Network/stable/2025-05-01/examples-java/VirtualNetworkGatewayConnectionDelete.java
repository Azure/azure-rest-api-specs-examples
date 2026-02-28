
/**
 * Samples for VirtualNetworkGatewayConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VirtualNetworkGatewayConnectionDelete.json
     */
    /**
     * Sample code: DeleteVirtualNetworkGatewayConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualNetworkGatewayConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayConnections().delete("rg1", "conn1",
            com.azure.core.util.Context.NONE);
    }
}
