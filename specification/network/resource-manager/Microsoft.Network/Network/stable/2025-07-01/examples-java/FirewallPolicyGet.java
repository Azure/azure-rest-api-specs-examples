
/**
 * Samples for FirewallPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyGet.json
     */
    /**
     * Sample code: Get FirewallPolicy.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getFirewallPolicy(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicies().getByResourceGroupWithResponse("rg1", "firewallPolicy", null,
            com.azure.core.util.Context.NONE);
    }
}
