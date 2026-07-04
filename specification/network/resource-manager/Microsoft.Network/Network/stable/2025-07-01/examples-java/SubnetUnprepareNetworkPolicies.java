
import com.azure.resourcemanager.network.models.UnprepareNetworkPoliciesRequest;

/**
 * Samples for Subnets UnprepareNetworkPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/SubnetUnprepareNetworkPolicies.json
     */
    /**
     * Sample code: Unprepare Network Policies.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void unprepareNetworkPolicies(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getSubnets().unprepareNetworkPolicies("rg1", "test-vnet", "subnet1",
            new UnprepareNetworkPoliciesRequest().withServiceName("Microsoft.Sql/managedInstances"),
            com.azure.core.util.Context.NONE);
    }
}
