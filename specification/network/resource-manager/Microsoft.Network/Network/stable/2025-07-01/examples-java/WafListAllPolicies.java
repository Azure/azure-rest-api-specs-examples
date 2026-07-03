
/**
 * Samples for WebApplicationFirewallPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WafListAllPolicies.json
     */
    /**
     * Sample code: Lists all WAF policies in a subscription.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listsAllWAFPoliciesInASubscription(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebApplicationFirewallPolicies().list(com.azure.core.util.Context.NONE);
    }
}
