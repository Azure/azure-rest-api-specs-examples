
/**
 * Samples for FirewallPolicyDrafts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyDraftGet.json
     */
    /**
     * Sample code: get firewall policy draft.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFirewallPolicyDraft(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyDrafts().getWithResponse("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
