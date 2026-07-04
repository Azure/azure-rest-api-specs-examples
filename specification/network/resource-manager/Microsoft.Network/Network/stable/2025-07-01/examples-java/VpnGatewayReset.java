
/**
 * Samples for VpnGateways Reset.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayReset.json
     */
    /**
     * Sample code: ResetVpnGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void resetVpnGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().reset("rg1", "vpngw", null, com.azure.core.util.Context.NONE);
    }
}
