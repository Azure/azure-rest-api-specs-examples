
/**
 * Samples for LoadBalancers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/LoadBalancerGetInboundNatRulePortMapping.json
     */
    /**
     * Sample code: Get load balancer with inbound NAT rule port mapping.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        getLoadBalancerWithInboundNATRulePortMapping(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().getByResourceGroupWithResponse("rg1", "lb", null, null,
            com.azure.core.util.Context.NONE);
    }
}
