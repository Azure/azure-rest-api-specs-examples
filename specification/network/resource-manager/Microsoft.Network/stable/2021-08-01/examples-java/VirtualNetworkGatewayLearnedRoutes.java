import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways GetLearnedRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewayLearnedRoutes.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayLearnedRoutes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayLearnedRoutes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .getLearnedRoutes("rg1", "vpngw", Context.NONE);
    }
}
