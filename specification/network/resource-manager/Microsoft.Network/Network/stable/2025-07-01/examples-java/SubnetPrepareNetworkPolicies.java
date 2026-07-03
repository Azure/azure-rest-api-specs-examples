
import com.azure.resourcemanager.network.models.PrepareNetworkPoliciesRequest;

/**
 * Samples for Subnets PrepareNetworkPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetPrepareNetworkPolicies.json
     */
    /**
     * Sample code: Prepare Network Policies.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void prepareNetworkPolicies(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().prepareNetworkPolicies("rg1", "test-vnet", "subnet1",
            new PrepareNetworkPoliciesRequest().withServiceName("Microsoft.Sql/managedInstances"),
            com.azure.core.util.Context.NONE);
    }
}
