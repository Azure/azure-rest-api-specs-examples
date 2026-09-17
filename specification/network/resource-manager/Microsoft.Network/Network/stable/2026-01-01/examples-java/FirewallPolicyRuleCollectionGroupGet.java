
/**
 * Samples for FirewallPolicyRuleCollectionGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/FirewallPolicyRuleCollectionGroupGet.json
     */
    /**
     * Sample code: Get FirewallPolicyRuleCollectionGroup.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFirewallPolicyRuleCollectionGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().getWithResponse("rg1", "firewallPolicy",
            "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
