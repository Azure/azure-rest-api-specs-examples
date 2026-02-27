
/**
 * Samples for VpnLinkConnections GetAllSharedKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * VpnSiteLinkConnectionSharedKeysGet.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionSharedKeysGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionSharedKeysGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnLinkConnections().getAllSharedKeys("rg1", "gateway1",
            "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
