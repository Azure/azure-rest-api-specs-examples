
/**
 * Samples for VpnLinkConnections ResetConnection.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VpnSiteLinkConnectionReset.
     * json
     */
    /**
     * Sample code: ResetVpnLinkConnection.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resetVpnLinkConnection(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnLinkConnections().resetConnection("rg1", "gateway1",
            "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
