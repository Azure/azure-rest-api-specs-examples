
/**
 * Samples for VpnSiteLinks ListByVpnSite.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VpnSiteLinkListByVpnSite.json
     */
    /**
     * Sample code: VpnSiteLinkListByVpnSite.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkListByVpnSite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnSiteLinks().listByVpnSite("rg1", "vpnSite1",
            com.azure.core.util.Context.NONE);
    }
}
