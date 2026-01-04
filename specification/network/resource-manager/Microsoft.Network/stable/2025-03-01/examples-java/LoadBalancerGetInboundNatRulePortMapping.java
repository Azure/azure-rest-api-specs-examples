
/**
 * Samples for LoadBalancers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * LoadBalancerGetInboundNatRulePortMapping.json
     */
    /**
     * Sample code: Get load balancer with inbound NAT rule port mapping.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getLoadBalancerWithInboundNATRulePortMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().getByResourceGroupWithResponse("rg1", "lb", null,
            com.azure.core.util.Context.NONE);
    }
}
