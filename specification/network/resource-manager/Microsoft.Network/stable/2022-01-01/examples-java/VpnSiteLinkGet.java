import com.azure.core.util.Context;

/** Samples for VpnSiteLinks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/VpnSiteLinkGet.json
     */
    /**
     * Sample code: VpnSiteGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnSiteLinks()
            .getWithResponse("rg1", "vpnSite1", "vpnSiteLink1", Context.NONE);
    }
}
