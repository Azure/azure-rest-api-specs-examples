
/**
 * Samples for WebApplicationFirewallPolicies ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/WafListPolicies.json
     */
    /**
     * Sample code: Lists all WAF policies in a resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listsAllWAFPoliciesInAResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getWebApplicationFirewallPolicies().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
