
/**
 * Samples for FirewallPolicyRuleCollectionGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * FirewallPolicyRuleCollectionGroupWithIpGroupsList.json
     */
    /**
     * Sample code: List all FirewallPolicyRuleCollectionGroups with IpGroups for a given FirewallPolicy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllFirewallPolicyRuleCollectionGroupsWithIpGroupsForAGivenFirewallPolicy(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroups().list("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
