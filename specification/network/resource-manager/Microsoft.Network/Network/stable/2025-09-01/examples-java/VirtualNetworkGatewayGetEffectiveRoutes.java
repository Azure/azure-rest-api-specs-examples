
/**
 * Samples for VirtualNetworkGateways GetEffectiveRoutes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/VirtualNetworkGatewayGetEffectiveRoutes.json
     */
    /**
     * Sample code: GetVirtualNetworkGatewayEffectiveRoutes.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getVirtualNetworkGatewayEffectiveRoutes(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().getEffectiveRoutes("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
