
/**
 * Samples for FirewallPolicyDrafts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/FirewallPolicyDraftDelete.
     * json
     */
    /**
     * Sample code: delete firewall policy draft.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFirewallPolicyDraft(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyDrafts().deleteWithResponse("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
