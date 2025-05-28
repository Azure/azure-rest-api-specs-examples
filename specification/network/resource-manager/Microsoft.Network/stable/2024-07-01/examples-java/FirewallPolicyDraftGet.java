
/**
 * Samples for FirewallPolicyDrafts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/FirewallPolicyDraftGet.json
     */
    /**
     * Sample code: get firewall policy draft.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicyDraft(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicyDrafts().getWithResponse("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
