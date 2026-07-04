
/**
 * Samples for P2SVpnGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/P2SVpnGatewayGet.json
     */
    /**
     * Sample code: P2SVpnGatewayGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void p2SVpnGatewayGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().getByResourceGroupWithResponse("rg1", "p2sVpnGateway1",
            com.azure.core.util.Context.NONE);
    }
}
