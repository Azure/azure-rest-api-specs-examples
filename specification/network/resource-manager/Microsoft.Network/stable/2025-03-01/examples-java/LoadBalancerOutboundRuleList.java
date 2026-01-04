
/**
 * Samples for LoadBalancerOutboundRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerOutboundRuleList.
     * json
     */
    /**
     * Sample code: LoadBalancerOutboundRuleList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void loadBalancerOutboundRuleList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancerOutboundRules().list("testrg", "lb1",
            com.azure.core.util.Context.NONE);
    }
}
