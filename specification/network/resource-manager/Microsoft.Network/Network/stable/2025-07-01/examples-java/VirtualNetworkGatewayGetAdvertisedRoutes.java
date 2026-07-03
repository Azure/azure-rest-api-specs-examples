
/**
 * Samples for VirtualNetworkGateways GetAdvertisedRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayGetAdvertisedRoutes.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayAdvertisedRoutes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayAdvertisedRoutes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getAdvertisedRoutes("rg1", "vpngw", "test",
            com.azure.core.util.Context.NONE);
    }
}
