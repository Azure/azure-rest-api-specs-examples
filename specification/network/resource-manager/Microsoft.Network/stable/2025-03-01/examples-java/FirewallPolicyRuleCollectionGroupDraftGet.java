
/**
 * Samples for FirewallPolicyRuleCollectionGroupDrafts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * FirewallPolicyRuleCollectionGroupDraftGet.json
     */
    /**
     * Sample code: get rule collection group draft.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRuleCollectionGroupDraft(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroupDrafts().getWithResponse("rg1",
            "firewallPolicy", "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
