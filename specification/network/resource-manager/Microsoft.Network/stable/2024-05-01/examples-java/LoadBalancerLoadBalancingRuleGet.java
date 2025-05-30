
/**
 * Samples for LoadBalancerLoadBalancingRules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/
     * LoadBalancerLoadBalancingRuleGet.json
     */
    /**
     * Sample code: LoadBalancerLoadBalancingRuleGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerLoadBalancingRuleGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerLoadBalancingRules().getWithResponse("testrg", "lb1",
            "rule1", com.azure.core.util.Context.NONE);
    }
}
