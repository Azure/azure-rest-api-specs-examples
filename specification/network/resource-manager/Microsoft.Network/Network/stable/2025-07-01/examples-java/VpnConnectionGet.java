
/**
 * Samples for VpnConnections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnConnectionGet.json
     */
    /**
     * Sample code: VpnConnectionGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnConnectionGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnConnections().getWithResponse("rg1", "gateway1", "vpnConnection1",
            com.azure.core.util.Context.NONE);
    }
}
