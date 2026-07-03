
/**
 * Samples for VpnGateways GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VpnGatewayGet.json
     */
    /**
     * Sample code: VpnGatewayGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void vpnGatewayGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVpnGateways().getByResourceGroupWithResponse("rg1", "gateway1",
            com.azure.core.util.Context.NONE);
    }
}
