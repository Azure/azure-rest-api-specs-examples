import com.azure.core.util.Context;

/** Samples for VirtualNetworkGatewayNatRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewayNatRuleDelete.json
     */
    /**
     * Sample code: VirtualNetworkGatewayNatRuleDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualNetworkGatewayNatRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGatewayNatRules()
            .delete("rg1", "gateway1", "natRule1", Context.NONE);
    }
}
