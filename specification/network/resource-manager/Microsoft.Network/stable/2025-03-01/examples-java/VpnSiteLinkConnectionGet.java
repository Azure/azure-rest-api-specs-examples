
/**
 * Samples for VpnSiteLinkConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VpnSiteLinkConnectionGet.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void vpnSiteLinkConnectionGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVpnSiteLinkConnections().getWithResponse("rg1", "gateway1",
            "vpnConnection1", "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
