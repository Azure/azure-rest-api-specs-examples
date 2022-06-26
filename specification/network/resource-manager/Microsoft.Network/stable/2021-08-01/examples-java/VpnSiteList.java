import com.azure.core.util.Context;

/** Samples for VpnSites List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/VpnSiteList.json
     */
    /**
     * Sample code: VpnSiteList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnSites().list(Context.NONE);
    }
}
