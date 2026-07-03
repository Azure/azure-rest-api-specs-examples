
/**
 * Samples for P2SVpnGateways GetP2SVpnConnectionHealth.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/P2SVpnGatewayGetConnectionHealth.json
     */
    /**
     * Sample code: P2SVpnGatewayGetConnectionHealth.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void p2SVpnGatewayGetConnectionHealth(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().getP2SVpnConnectionHealth("rg1", "p2sVpnGateway1",
            com.azure.core.util.Context.NONE);
    }
}
