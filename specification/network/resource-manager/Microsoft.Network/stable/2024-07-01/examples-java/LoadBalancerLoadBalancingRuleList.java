
/**
 * Samples for LoadBalancerLoadBalancingRules List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/
     * LoadBalancerLoadBalancingRuleList.json
     */
    /**
     * Sample code: LoadBalancerLoadBalancingRuleList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerLoadBalancingRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerLoadBalancingRules().list("testrg", "lb1",
            com.azure.core.util.Context.NONE);
    }
}
