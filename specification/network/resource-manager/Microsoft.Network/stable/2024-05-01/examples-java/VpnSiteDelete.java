
/**
 * Samples for VpnSites Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VpnSiteDelete.json
     */
    /**
     * Sample code: VpnSiteDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnSites().delete("rg1", "vpnSite1",
            com.azure.core.util.Context.NONE);
    }
}
