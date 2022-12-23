import com.azure.core.util.Context;

/** Samples for VpnLinkConnections ListByVpnConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/VpnSiteLinkConnectionList.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVpnLinkConnections()
            .listByVpnConnection("rg1", "gateway1", "vpnConnection1", Context.NONE);
    }
}
