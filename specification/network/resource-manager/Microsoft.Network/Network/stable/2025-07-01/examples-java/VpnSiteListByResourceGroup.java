
/**
 * Samples for VpnSites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteListByResourceGroup.json
     */
    /**
     * Sample code: VpnSiteListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSites().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
