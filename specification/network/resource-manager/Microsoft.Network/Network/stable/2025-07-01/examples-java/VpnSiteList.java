
/**
 * Samples for VpnSites List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteList.json
     */
    /**
     * Sample code: VpnSiteList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSites().list(com.azure.core.util.Context.NONE);
    }
}
