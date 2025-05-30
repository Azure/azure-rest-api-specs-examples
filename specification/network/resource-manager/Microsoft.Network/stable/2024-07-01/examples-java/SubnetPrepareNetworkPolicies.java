
import com.azure.resourcemanager.network.models.PrepareNetworkPoliciesRequest;

/**
 * Samples for Subnets PrepareNetworkPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/SubnetPrepareNetworkPolicies.
     * json
     */
    /**
     * Sample code: Prepare Network Policies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void prepareNetworkPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubnets().prepareNetworkPolicies("rg1", "test-vnet", "subnet1",
            new PrepareNetworkPoliciesRequest().withServiceName("Microsoft.Sql/managedInstances"),
            com.azure.core.util.Context.NONE);
    }
}
