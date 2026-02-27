
/**
 * Samples for FirewallPolicyRuleCollectionGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * FirewallPolicyNatRuleCollectionGroupGet.json
     */
    /**
     * Sample code: Get FirewallPolicyNatRuleCollectionGroup.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicyNatRuleCollectionGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroups().getWithResponse("rg1",
            "firewallPolicy", "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
