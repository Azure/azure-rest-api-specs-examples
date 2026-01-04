
/**
 * Samples for FirewallPolicyRuleCollectionGroupDrafts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * FirewallPolicyRuleCollectionGroupDraftDelete.json
     */
    /**
     * Sample code: delete firewall rule collection group draft.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFirewallRuleCollectionGroupDraft(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyRuleCollectionGroupDrafts()
            .deleteWithResponse("rg1", "firewallPolicy", "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
