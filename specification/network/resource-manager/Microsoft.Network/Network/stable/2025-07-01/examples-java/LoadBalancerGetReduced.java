
import com.azure.resourcemanager.network.models.LoadBalancerDetailLevel;

/**
 * Samples for LoadBalancers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerGetReduced.json
     */
    /**
     * Sample code: Get load balancer reduced.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getLoadBalancerReduced(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().getByResourceGroupWithResponse("rg-name", "lb-name", null,
            LoadBalancerDetailLevel.REDUCED, com.azure.core.util.Context.NONE);
    }
}
