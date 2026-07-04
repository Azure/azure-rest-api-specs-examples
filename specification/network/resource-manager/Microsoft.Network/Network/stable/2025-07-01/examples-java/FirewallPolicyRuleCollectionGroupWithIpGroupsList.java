
/**
 * Samples for FirewallPolicyRuleCollectionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupWithIpGroupsList.json
     */
    /**
     * Sample code: List all FirewallPolicyRuleCollectionGroups with IpGroups for a given FirewallPolicy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listAllFirewallPolicyRuleCollectionGroupsWithIpGroupsForAGivenFirewallPolicy(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroups().list("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
