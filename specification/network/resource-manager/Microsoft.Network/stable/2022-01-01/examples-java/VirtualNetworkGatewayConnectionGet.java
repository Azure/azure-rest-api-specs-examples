import com.azure.core.util.Context;

/** Samples for VirtualNetworkGatewayConnections GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VirtualNetworkGatewayConnectionGet.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayConnection.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGatewayConnections()
            .getByResourceGroupWithResponse("rg1", "connS2S", Context.NONE);
    }
}
