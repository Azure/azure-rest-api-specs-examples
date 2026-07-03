
/**
 * Samples for FirewallPolicyRuleCollectionGroupDrafts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyRuleCollectionGroupDraftGet.json
     */
    /**
     * Sample code: get rule collection group draft.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getRuleCollectionGroupDraft(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyRuleCollectionGroupDrafts().getWithResponse("rg1", "firewallPolicy",
            "ruleCollectionGroup1", com.azure.core.util.Context.NONE);
    }
}
