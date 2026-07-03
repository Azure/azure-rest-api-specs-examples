
/**
 * Samples for FirewallPolicyDeployments Deploy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyDraftDeploy.json
     */
    /**
     * Sample code: deploy firewall policy draft.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deployFirewallPolicyDraft(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicyDeployments().deploy("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
