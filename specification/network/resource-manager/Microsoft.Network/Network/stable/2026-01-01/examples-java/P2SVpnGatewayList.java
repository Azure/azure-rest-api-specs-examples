
/**
 * Samples for P2SVpnGateways List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/P2SVpnGatewayList.json
     */
    /**
     * Sample code: P2SVpnGatewayListBySubscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void p2SVpnGatewayListBySubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().list(com.azure.core.util.Context.NONE);
    }
}
