import com.azure.core.util.Context;

/** Samples for WebApplicationFirewallPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/WafListAllPolicies.json
     */
    /**
     * Sample code: Lists all WAF policies in a subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllWAFPoliciesInASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getWebApplicationFirewallPolicies().list(Context.NONE);
    }
}
