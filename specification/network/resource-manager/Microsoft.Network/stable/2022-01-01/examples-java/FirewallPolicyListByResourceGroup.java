import com.azure.core.util.Context;

/** Samples for FirewallPolicies ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/FirewallPolicyListByResourceGroup.json
     */
    /**
     * Sample code: List all Firewall Policies for a given resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllFirewallPoliciesForAGivenResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getFirewallPolicies().listByResourceGroup("rg1", Context.NONE);
    }
}
