
/**
 * Samples for FirewallPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyDelete.json
     */
    /**
     * Sample code: Delete Firewall Policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFirewallPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicies().delete("rg1", "firewallPolicy",
            com.azure.core.util.Context.NONE);
    }
}
