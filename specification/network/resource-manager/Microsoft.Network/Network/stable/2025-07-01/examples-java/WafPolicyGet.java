
/**
 * Samples for WebApplicationFirewallPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/WafPolicyGet.json
     */
    /**
     * Sample code: Gets a WAF policy within a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getsAWAFPolicyWithinAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebApplicationFirewallPolicies().getByResourceGroupWithResponse("rg1", "Policy1",
            com.azure.core.util.Context.NONE);
    }
}
