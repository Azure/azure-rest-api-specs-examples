
import com.azure.resourcemanager.network.models.QueryInboundNatRulePortMappingRequest;

/**
 * Samples for LoadBalancers ListInboundNatRulePortMappings.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-03-01/examples/
     * QueryInboundNatRulePortMapping.json
     */
    /**
     * Sample code: Query inbound NAT rule port mapping.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void queryInboundNATRulePortMapping(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().listInboundNatRulePortMappings("rg1", "lb1",
            "bp1", new QueryInboundNatRulePortMappingRequest().withIpAddress("10.0.0.4"),
            com.azure.core.util.Context.NONE);
    }
}
