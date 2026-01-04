
/**
 * Samples for LoadBalancerLoadBalancingRules Health.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerHealth.json
     */
    /**
     * Sample code: Query load balancing rule health.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queryLoadBalancingRuleHealth(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerLoadBalancingRules().health("rg1", "lb1", "rulelb",
            com.azure.core.util.Context.NONE);
    }
}
