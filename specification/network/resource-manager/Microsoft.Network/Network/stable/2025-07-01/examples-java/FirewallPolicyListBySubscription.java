
/**
 * Samples for FirewallPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/FirewallPolicyListBySubscription.json
     */
    /**
     * Sample code: List all Firewall Policies for a given subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        listAllFirewallPoliciesForAGivenSubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getFirewallPolicies().list(com.azure.core.util.Context.NONE);
    }
}
