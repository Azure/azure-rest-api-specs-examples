
/**
 * Samples for LoadBalancers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerGet.json
     */
    /**
     * Sample code: Get load balancer.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getLoadBalancer(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().getByResourceGroupWithResponse("rg1", "lb", null, null,
            com.azure.core.util.Context.NONE);
    }
}
