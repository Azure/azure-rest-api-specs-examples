
/**
 * Samples for VpnSites GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteGet.json
     */
    /**
     * Sample code: VpnSiteGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSites().getByResourceGroupWithResponse("rg1", "vpnSite1",
            com.azure.core.util.Context.NONE);
    }
}
