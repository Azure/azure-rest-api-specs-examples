
/**
 * Samples for VpnLinkConnections ListDefaultSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkConnectionDefaultSharedKeyList.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionDefaultSharedKeyList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        vpnSiteLinkConnectionDefaultSharedKeyList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnLinkConnections().listDefaultSharedKeyWithResponse("rg1", "gateway1",
            "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
