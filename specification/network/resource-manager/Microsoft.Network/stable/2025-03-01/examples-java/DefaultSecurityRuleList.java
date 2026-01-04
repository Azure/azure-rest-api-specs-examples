
/**
 * Samples for DefaultSecurityRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/DefaultSecurityRuleList.json
     */
    /**
     * Sample code: DefaultSecurityRuleList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void defaultSecurityRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getDefaultSecurityRules().list("testrg", "nsg1",
            com.azure.core.util.Context.NONE);
    }
}
