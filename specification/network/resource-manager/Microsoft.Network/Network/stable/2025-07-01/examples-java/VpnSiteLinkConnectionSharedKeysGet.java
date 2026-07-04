
/**
 * Samples for VpnLinkConnections GetAllSharedKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkConnectionSharedKeysGet.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionSharedKeysGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteLinkConnectionSharedKeysGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnLinkConnections().getAllSharedKeys("rg1", "gateway1", "vpnConnection1",
            "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
