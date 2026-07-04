
/**
 * Samples for FirewallPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyDelete.json
     */
    /**
     * Sample code: Delete Firewall Policy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteFirewallPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicies().delete("rg1", "firewallPolicy", com.azure.core.util.Context.NONE);
    }
}
