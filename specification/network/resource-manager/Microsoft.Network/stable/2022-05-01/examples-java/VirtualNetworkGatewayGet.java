import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewayGet.json
     */
    /**
     * Sample code: GetVirtualNetworkGateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .getByResourceGroupWithResponse("rg1", "vpngw", Context.NONE);
    }
}
