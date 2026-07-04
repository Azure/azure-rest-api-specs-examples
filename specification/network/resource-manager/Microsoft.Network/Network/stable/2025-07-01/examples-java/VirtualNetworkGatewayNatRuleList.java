
/**
 * Samples for VirtualNetworkGatewayNatRules ListByVirtualNetworkGateway.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayNatRuleList.json
     */
    /**
     * Sample code: VirtualNetworkGatewayNatRuleList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualNetworkGatewayNatRuleList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayNatRules().listByVirtualNetworkGateway("rg1", "gateway1",
            com.azure.core.util.Context.NONE);
    }
}
