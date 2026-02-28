
/**
 * Samples for FirewallPolicies GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyGet.json
     */
    /**
     * Sample code: Get FirewallPolicy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicies().getByResourceGroupWithResponse("rg1",
            "firewallPolicy", null, com.azure.core.util.Context.NONE);
    }
}
