
/**
 * Samples for FirewallPolicyRuleCollectionGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithIpGroupsGet.json
     */
    /**
     * Sample code: Get FirewallPolicyRuleCollectionGroup With IpGroups.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getFirewallPolicyRuleCollectionGroupWithIpGroups(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().getWithResponse("rg1", "firewallPolicy",
            "ruleGroup1", com.azure.core.util.Context.NONE);
    }
}
