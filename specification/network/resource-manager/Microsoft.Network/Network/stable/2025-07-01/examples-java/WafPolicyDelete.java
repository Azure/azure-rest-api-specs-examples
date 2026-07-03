
/**
 * Samples for WebApplicationFirewallPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WafPolicyDelete.json
     */
    /**
     * Sample code: Deletes a WAF policy within a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deletesAWAFPolicyWithinAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebApplicationFirewallPolicies().delete("rg1", "Policy1",
            com.azure.core.util.Context.NONE);
    }
}
