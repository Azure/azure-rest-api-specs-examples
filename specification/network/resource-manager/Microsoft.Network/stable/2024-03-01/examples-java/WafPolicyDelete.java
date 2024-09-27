
/**
 * Samples for WebApplicationFirewallPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/WafPolicyDelete.json
     */
    /**
     * Sample code: Deletes a WAF policy within a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesAWAFPolicyWithinAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getWebApplicationFirewallPolicies().delete("rg1", "Policy1",
            com.azure.core.util.Context.NONE);
    }
}
