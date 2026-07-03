
/**
 * Samples for VpnGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayList.json
     */
    /**
     * Sample code: VpnGatewayListBySubscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnGatewayListBySubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().list(com.azure.core.util.Context.NONE);
    }
}
