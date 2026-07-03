
/**
 * Samples for VpnLinkConnections GetDefaultSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkConnectionDefaultSharedKeyGet.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionDefaultSharedKeyGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        vpnSiteLinkConnectionDefaultSharedKeyGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnLinkConnections().getDefaultSharedKeyWithResponse("rg1", "gateway1",
            "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
