import com.azure.core.util.Context;

/** Samples for VirtualNetworkGateways GetAdvertisedRoutes. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VirtualNetworkGatewayGetAdvertisedRoutes.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayAdvertisedRoutes.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualNetworkGatewayAdvertisedRoutes(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .getAdvertisedRoutes("rg1", "vpngw", "test", Context.NONE);
    }
}
