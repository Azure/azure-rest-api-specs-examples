
/**
 * Samples for P2SVpnGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/P2SVpnGatewayListByResourceGroup.json
     */
    /**
     * Sample code: P2SVpnGatewayListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void p2SVpnGatewayListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getP2SVpnGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
