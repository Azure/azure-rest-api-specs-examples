
/**
 * Samples for P2SVpnGateways Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/P2SVpnGatewayDelete.json
     */
    /**
     * Sample code: P2SVpnGatewayDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void p2SVpnGatewayDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().delete("rg1", "p2sVpnGateway1", com.azure.core.util.Context.NONE);
    }
}
