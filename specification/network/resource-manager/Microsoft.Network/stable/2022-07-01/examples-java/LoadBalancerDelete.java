import com.azure.core.util.Context;

/** Samples for LoadBalancers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/LoadBalancerDelete.json
     */
    /**
     * Sample code: Delete load balancer.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteLoadBalancer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().delete("rg1", "lb", Context.NONE);
    }
}
