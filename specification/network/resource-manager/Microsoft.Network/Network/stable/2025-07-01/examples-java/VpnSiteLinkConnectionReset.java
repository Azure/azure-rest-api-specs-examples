
/**
 * Samples for VpnLinkConnections ResetConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkConnectionReset.json
     */
    /**
     * Sample code: ResetVpnLinkConnection.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void resetVpnLinkConnection(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnLinkConnections().resetConnection("rg1", "gateway1", "vpnConnection1",
            "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
