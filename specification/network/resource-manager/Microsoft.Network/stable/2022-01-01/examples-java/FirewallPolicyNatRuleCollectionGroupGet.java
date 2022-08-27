import com.azure.core.util.Context;

/** Samples for FirewallPolicyRuleCollectionGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/FirewallPolicyNatRuleCollectionGroupGet.json
     */
    /**
     * Sample code: Get FirewallPolicyNatRuleCollectionGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicyNatRuleCollectionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicyRuleCollectionGroups()
            .getWithResponse("rg1", "firewallPolicy", "ruleCollectionGroup1", Context.NONE);
    }
}
