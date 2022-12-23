import com.azure.core.util.Context;

/** Samples for LoadBalancerBackendAddressPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/LoadBalancerBackendAddressPoolGet.json
     */
    /**
     * Sample code: LoadBalancerBackendAddressPoolGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerBackendAddressPoolGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancerBackendAddressPools()
            .getWithResponse("testrg", "lb", "backend", Context.NONE);
    }
}
