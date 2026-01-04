
/**
 * Samples for VirtualNetworkGatewayNatRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * VirtualNetworkGatewayNatRuleGet.json
     */
    /**
     * Sample code: VirtualNetworkGatewayNatRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewayNatRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGatewayNatRules().getWithResponse("rg1", "gateway1",
            "natRule1", com.azure.core.util.Context.NONE);
    }
}
