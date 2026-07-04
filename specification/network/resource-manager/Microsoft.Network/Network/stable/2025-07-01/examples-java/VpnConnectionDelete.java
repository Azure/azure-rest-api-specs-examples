
/**
 * Samples for VpnConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnConnectionDelete.json
     */
    /**
     * Sample code: VpnConnectionDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnConnectionDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnConnections().delete("rg1", "gateway1", "vpnConnection1",
            com.azure.core.util.Context.NONE);
    }
}
