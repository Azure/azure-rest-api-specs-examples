
/**
 * Samples for P2SVpnGateways Reset.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/P2SVpnGatewayReset.json
     */
    /**
     * Sample code: ResetP2SVpnGateway.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void resetP2SVpnGateway(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().reset("rg1", "p2sVpnGateway1", com.azure.core.util.Context.NONE);
    }
}
