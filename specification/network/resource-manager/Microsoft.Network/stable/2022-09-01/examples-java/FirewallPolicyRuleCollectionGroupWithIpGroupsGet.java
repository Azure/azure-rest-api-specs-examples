/** Samples for FirewallPolicyRuleCollectionGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-09-01/examples/FirewallPolicyRuleCollectionGroupWithIpGroupsGet.json
     */
    /**
     * Sample code: Get FirewallPolicyRuleCollectionGroup With IpGroups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicyRuleCollectionGroupWithIpGroups(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyRuleCollectionGroups()
            .getWithResponse("rg1", "firewallPolicy", "ruleGroup1", com.azure.core.util.Context.NONE);
    }
}
