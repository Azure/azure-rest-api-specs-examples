
/**
 * Samples for FirewallPolicyRuleCollectionGroupDrafts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupDraftDelete.json
     */
    /**
     * Sample code: delete firewall rule collection group draft.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        deleteFirewallRuleCollectionGroupDraft(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroupDrafts().deleteWithResponse("rg1", "firewallPolicy",
            "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
