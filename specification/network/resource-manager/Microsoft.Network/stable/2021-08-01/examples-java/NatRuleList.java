import com.azure.core.util.Context;

/** Samples for NatRules ListByVpnGateway. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NatRuleList.json
     */
    /**
     * Sample code: NatRuleList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void natRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatRules().listByVpnGateway("rg1", "gateway1", Context.NONE);
    }
}
