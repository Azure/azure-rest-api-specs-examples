
/**
 * Samples for VirtualNetworkGateways ResetVpnClientSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayResetVpnClientSharedKey.json
     */
    /**
     * Sample code: ResetVpnClientSharedKey.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void resetVpnClientSharedKey(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGateways().resetVpnClientSharedKey("rg1", "vpngw",
            com.azure.core.util.Context.NONE);
    }
}
