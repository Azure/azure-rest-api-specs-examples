
/**
 * Samples for FirewallPolicyDeployments Deploy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyDraftDeploy.
     * json
     */
    /**
     * Sample code: deploy firewall policy draft.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deployFirewallPolicyDraft(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyDeployments().deploy("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
