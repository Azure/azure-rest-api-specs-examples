
/**
 * Samples for VpnGateways ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayListByResourceGroup.json
     */
    /**
     * Sample code: VpnGatewayListByResourceGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnGatewayListByResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
