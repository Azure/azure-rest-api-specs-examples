import com.azure.core.util.Context;

/** Samples for FirewallPolicies GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/FirewallPolicyGet.json
     */
    /**
     * Sample code: Get FirewallPolicy.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFirewallPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getFirewallPolicies()
            .getByResourceGroupWithResponse("rg1", "firewallPolicy", null, Context.NONE);
    }
}
