
/**
 * Samples for WebApplicationFirewallPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/WafPolicyGetBasic.json
     */
    /**
     * Sample code: Gets a Basic tier WAF policy within a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getsABasicTierWAFPolicyWithinAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebApplicationFirewallPolicies().getByResourceGroupWithResponse("rg1", "Policy1",
            com.azure.core.util.Context.NONE);
    }
}
