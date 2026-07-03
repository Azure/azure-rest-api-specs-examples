
/**
 * Samples for LoadBalancerLoadBalancingRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerLoadBalancingRuleList.json
     */
    /**
     * Sample code: LoadBalancerLoadBalancingRuleList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerLoadBalancingRuleList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerLoadBalancingRules().list("testrg", "lb1",
            com.azure.core.util.Context.NONE);
    }
}
