
import com.azure.resourcemanager.network.models.UnprepareNetworkPoliciesRequest;

/**
 * Samples for Subnets UnprepareNetworkPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * SubnetUnprepareNetworkPolicies.json
     */
    /**
     * Sample code: Unprepare Network Policies.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void unprepareNetworkPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getSubnets().unprepareNetworkPolicies("rg1", "test-vnet", "subnet1",
            new UnprepareNetworkPoliciesRequest().withServiceName("Microsoft.Sql/managedInstances"),
            com.azure.core.util.Context.NONE);
    }
}
