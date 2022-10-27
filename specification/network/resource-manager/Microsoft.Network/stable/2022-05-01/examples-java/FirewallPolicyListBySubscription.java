import com.azure.core.util.Context;

/** Samples for FirewallPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/FirewallPolicyListBySubscription.json
     */
    /**
     * Sample code: List all Firewall Policies for a given subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllFirewallPoliciesForAGivenSubscription(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicies().list(Context.NONE);
    }
}
