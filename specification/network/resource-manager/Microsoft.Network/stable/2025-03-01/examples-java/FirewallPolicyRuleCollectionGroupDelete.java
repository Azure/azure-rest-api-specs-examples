
/**
 * Samples for FirewallPolicyRuleCollectionGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * FirewallPolicyRuleCollectionGroupDelete.json
     */
    /**
     * Sample code: Delete FirewallPolicyRuleCollectionGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFirewallPolicyRuleCollectionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroups().delete("rg1",
            "firewallPolicy", "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
