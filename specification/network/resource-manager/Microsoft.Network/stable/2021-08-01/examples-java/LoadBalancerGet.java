import com.azure.core.util.Context;

/** Samples for LoadBalancers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/LoadBalancerGet.json
     */
    /**
     * Sample code: Get load balancer.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLoadBalancer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancers()
            .getByResourceGroupWithResponse("rg1", "lb", null, Context.NONE);
    }
}
