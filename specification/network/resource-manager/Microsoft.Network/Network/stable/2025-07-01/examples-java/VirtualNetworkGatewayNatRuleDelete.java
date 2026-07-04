
/**
 * Samples for VirtualNetworkGatewayNatRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkGatewayNatRuleDelete.json
     */
    /**
     * Sample code: VirtualNetworkGatewayNatRuleDelete.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void virtualNetworkGatewayNatRuleDelete(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkGatewayNatRules().delete("rg1", "gateway1", "natRule1",
            com.azure.core.util.Context.NONE);
    }
}
