
/**
 * Samples for VirtualNetworkGateways Reset.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayReset.json
     */
    /**
     * Sample code: ResetVirtualNetworkGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void resetVirtualNetworkGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().reset("rg1", "vpngw", null,
            com.azure.core.util.Context.NONE);
    }
}
