
/**
 * Samples for NatRules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NatRuleDelete.json
     */
    /**
     * Sample code: NatRuleDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void natRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNatRules().delete("rg1", "gateway1", "natRule1",
            com.azure.core.util.Context.NONE);
    }
}
