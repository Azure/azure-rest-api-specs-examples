import com.azure.core.util.Context;

/** Samples for VirtualNetworkGatewayConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VirtualNetworkGatewayConnectionDelete.json
     */
    /**
     * Sample code: DeleteVirtualNetworkGatewayConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteVirtualNetworkGatewayConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGatewayConnections()
            .delete("rg1", "conn1", Context.NONE);
    }
}
