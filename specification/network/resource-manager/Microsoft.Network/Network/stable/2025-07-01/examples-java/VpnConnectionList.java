
/**
 * Samples for VpnConnections ListByVpnGateway.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnConnectionList.json
     */
    /**
     * Sample code: VpnConnectionList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnConnectionList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnConnections().listByVpnGateway("rg1", "gateway1",
            com.azure.core.util.Context.NONE);
    }
}
