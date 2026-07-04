
/**
 * Samples for LoadBalancerOutboundRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerOutboundRuleList.json
     */
    /**
     * Sample code: LoadBalancerOutboundRuleList.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void loadBalancerOutboundRuleList(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerOutboundRules().list("testrg", "lb1", com.azure.core.util.Context.NONE);
    }
}
