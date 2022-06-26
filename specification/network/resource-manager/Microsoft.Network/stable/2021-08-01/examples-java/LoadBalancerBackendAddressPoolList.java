import com.azure.core.util.Context;

/** Samples for LoadBalancerBackendAddressPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LoadBalancerBackendAddressPoolList.json
     */
    /**
     * Sample code: LoadBalancerBackendAddressPoolList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerBackendAddressPoolList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancerBackendAddressPools()
            .list("testrg", "lb", Context.NONE);
    }
}
