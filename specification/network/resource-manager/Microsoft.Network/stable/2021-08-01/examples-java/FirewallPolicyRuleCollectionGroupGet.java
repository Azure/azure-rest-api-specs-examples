import com.azure.core.util.Context;

/** Samples for FirewallPolicyRuleCollectionGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicyRuleCollectionGroupGet.json
     */
    /**
     * Sample code: Get FirewallPolicyRuleCollectionGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicyRuleCollectionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyRuleCollectionGroups()
            .getWithResponse("rg1", "firewallPolicy", "ruleCollectionGroup1", Context.NONE);
    }
}
