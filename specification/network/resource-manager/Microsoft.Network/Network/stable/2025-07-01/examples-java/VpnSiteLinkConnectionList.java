
/**
 * Samples for VpnLinkConnections ListByVpnConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkConnectionList.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteLinkConnectionList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnLinkConnections().listByVpnConnection("rg1", "gateway1", "vpnConnection1",
            com.azure.core.util.Context.NONE);
    }
}
