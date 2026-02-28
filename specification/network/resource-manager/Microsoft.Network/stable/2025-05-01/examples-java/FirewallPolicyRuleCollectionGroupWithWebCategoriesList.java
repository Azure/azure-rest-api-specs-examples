
/**
 * Samples for FirewallPolicyRuleCollectionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * FirewallPolicyRuleCollectionGroupWithWebCategoriesList.json
     */
    /**
     * Sample code: List all FirewallPolicyRuleCollectionGroup With Web Categories.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllFirewallPolicyRuleCollectionGroupWithWebCategories(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroups().list("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
