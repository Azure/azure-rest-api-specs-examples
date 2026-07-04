
/**
 * Samples for FirewallPolicyRuleCollectionGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithWebCategoriesGet.json
     */
    /**
     * Sample code: Get FirewallPolicyRuleCollectionGroup With Web Categories.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFirewallPolicyRuleCollectionGroupWithWebCategories(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().getWithResponse("rg1", "firewallPolicy",
            "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
