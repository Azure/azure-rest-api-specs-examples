
/**
 * Samples for VpnLinkConnections GetDefaultSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VpnSiteLinkConnectionDefaultSharedKeyGet.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionDefaultSharedKeyGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionDefaultSharedKeyGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnLinkConnections().getDefaultSharedKeyWithResponse("rg1",
            "gateway1", "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
