
/**
 * Samples for VpnSites ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnSiteListByResourceGroup.
     * json
     */
    /**
     * Sample code: VpnSiteListByResourceGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnSites().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
