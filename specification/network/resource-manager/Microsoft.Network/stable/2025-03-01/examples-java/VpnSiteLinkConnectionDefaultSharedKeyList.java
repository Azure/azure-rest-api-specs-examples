
/**
 * Samples for VpnLinkConnections ListDefaultSharedKey.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VpnSiteLinkConnectionDefaultSharedKeyList.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionDefaultSharedKeyList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionDefaultSharedKeyList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnLinkConnections().listDefaultSharedKeyWithResponse("rg1",
            "gateway1", "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
