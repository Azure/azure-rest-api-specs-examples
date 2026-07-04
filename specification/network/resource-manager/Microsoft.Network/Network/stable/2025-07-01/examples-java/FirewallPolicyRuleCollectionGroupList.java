
/**
 * Samples for FirewallPolicyRuleCollectionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupList.json
     */
    /**
     * Sample code: List all FirewallPolicyRuleCollectionGroups for a given FirewallPolicy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllFirewallPolicyRuleCollectionGroupsForAGivenFirewallPolicy(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().list("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
