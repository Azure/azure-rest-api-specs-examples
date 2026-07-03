
/**
 * Samples for VpnSiteLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkGet.json
     */
    /**
     * Sample code: VpnSiteGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSiteLinks().getWithResponse("rg1", "vpnSite1", "vpnSiteLink1",
            com.azure.core.util.Context.NONE);
    }
}
