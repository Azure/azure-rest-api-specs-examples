
/**
 * Samples for LoadBalancerOutboundRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerOutboundRuleGet.json
     */
    /**
     * Sample code: LoadBalancerOutboundRuleGet.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerOutboundRuleGet(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerOutboundRules().getWithResponse("testrg", "lb1", "rule1",
            com.azure.core.util.Context.NONE);
    }
}
