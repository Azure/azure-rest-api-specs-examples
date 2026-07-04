
/**
 * Samples for VpnGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayDelete.json
     */
    /**
     * Sample code: VpnGatewayDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnGatewayDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().delete("rg1", "gateway1", com.azure.core.util.Context.NONE);
    }
}
