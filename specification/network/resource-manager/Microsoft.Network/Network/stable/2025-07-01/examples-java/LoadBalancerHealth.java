
/**
 * Samples for LoadBalancerLoadBalancingRules Health.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerHealth.json
     */
    /**
     * Sample code: Query load balancing rule health.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void queryLoadBalancingRuleHealth(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancerLoadBalancingRules().health("rg1", "lb1", "rulelb",
            com.azure.core.util.Context.NONE);
    }
}
