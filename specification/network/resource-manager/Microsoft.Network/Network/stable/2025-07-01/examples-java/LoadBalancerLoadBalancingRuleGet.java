
/**
 * Samples for LoadBalancerLoadBalancingRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerLoadBalancingRuleGet.json
     */
    /**
     * Sample code: LoadBalancerLoadBalancingRuleGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerLoadBalancingRuleGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerLoadBalancingRules().getWithResponse("testrg", "lb1", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
