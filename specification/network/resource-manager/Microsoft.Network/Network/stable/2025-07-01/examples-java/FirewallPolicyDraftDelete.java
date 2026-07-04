
/**
 * Samples for FirewallPolicyDrafts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyDraftDelete.json
     */
    /**
     * Sample code: delete firewall policy draft.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteFirewallPolicyDraft(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyDrafts().deleteWithResponse("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
