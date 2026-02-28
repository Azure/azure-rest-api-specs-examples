
/**
 * Samples for VpnSites GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnSiteGet.json
     */
    /**
     * Sample code: VpnSiteGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnSites().getByResourceGroupWithResponse("rg1", "vpnSite1",
            com.azure.core.util.Context.NONE);
    }
}
