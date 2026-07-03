
/**
 * Samples for VpnSiteLinkConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnSiteLinkConnectionGet.json
     */
    /**
     * Sample code: VpnSiteLinkConnectionGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnSiteLinkConnectionGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnSiteLinkConnections().getWithResponse("rg1", "gateway1", "vpnConnection1",
            "Connection-Link1", com.azure.core.util.Context.NONE);
    }
}
