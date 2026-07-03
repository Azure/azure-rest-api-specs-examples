
/**
 * Samples for VirtualNetworkGateways GetLearnedRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayLearnedRoutes.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayLearnedRoutes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getVirtualNetworkGatewayLearnedRoutes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getLearnedRoutes("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
