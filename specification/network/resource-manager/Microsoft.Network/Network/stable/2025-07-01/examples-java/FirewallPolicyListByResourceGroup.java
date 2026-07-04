
/**
 * Samples for FirewallPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyListByResourceGroup.json
     */
    /**
     * Sample code: List all Firewall Policies for a given resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllFirewallPoliciesForAGivenResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicies().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
