
/**
 * Samples for FirewallPolicyRuleCollectionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithWebCategoriesList.json
     */
    /**
     * Sample code: List all FirewallPolicyRuleCollectionGroup With Web Categories.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllFirewallPolicyRuleCollectionGroupWithWebCategories(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().list("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
