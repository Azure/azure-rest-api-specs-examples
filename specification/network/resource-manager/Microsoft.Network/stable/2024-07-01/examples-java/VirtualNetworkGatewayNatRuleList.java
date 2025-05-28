
/**
 * Samples for VirtualNetworkGatewayNatRules ListByVirtualNetworkGateway.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * VirtualNetworkGatewayNatRuleList.json
     */
    /**
     * Sample code: VirtualNetworkGatewayNatRuleList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewayNatRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayNatRules().listByVirtualNetworkGateway("rg1",
            "gateway1", com.azure.core.util.Context.NONE);
    }
}
