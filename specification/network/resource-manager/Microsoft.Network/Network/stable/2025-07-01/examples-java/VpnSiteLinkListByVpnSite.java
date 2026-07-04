
/**
 * Samples for VpnSiteLinks ListByVpnSite.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkListByVpnSite.json
     */
    /**
     * Sample code: VpnSiteLinkListByVpnSite.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteLinkListByVpnSite(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSiteLinks().listByVpnSite("rg1", "vpnSite1", com.azure.core.util.Context.NONE);
    }
}
