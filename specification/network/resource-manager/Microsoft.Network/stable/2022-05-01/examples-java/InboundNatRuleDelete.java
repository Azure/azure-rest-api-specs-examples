import com.azure.core.util.Context;

/** Samples for InboundNatRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/InboundNatRuleDelete.json
     */
    /**
     * Sample code: InboundNatRuleDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void inboundNatRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getInboundNatRules()
            .delete("testrg", "lb1", "natRule1.1", Context.NONE);
    }
}
