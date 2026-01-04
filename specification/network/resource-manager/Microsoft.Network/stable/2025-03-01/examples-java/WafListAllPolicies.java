
/**
 * Samples for WebApplicationFirewallPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/WafListAllPolicies.json
     */
    /**
     * Sample code: Lists all WAF policies in a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllWAFPoliciesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getWebApplicationFirewallPolicies()
            .list(com.azure.core.util.Context.NONE);
    }
}
