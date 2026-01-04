
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.FrontendIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancerInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.InboundNatPool;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.LoadBalancerSku;
import com.azure.resourcemanager.network.models.LoadBalancerSkuName;
import com.azure.resourcemanager.network.models.TransportProtocol;
import java.util.Arrays;

/**
 * Samples for LoadBalancers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/
     * LoadBalancerCreateWithInboundNatPool.json
     */
    /**
     * Sample code: Create load balancer with inbound nat pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createLoadBalancerWithInboundNatPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().createOrUpdate("rg1", "lb",
            new LoadBalancerInner().withLocation("eastus")
                .withSku(new LoadBalancerSku().withName(LoadBalancerSkuName.STANDARD))
                .withFrontendIpConfigurations(Arrays.asList(new FrontendIpConfigurationInner().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test")
                    .withName("test").withZones(Arrays.asList())
                    .withPrivateIpAllocationMethod(IpAllocationMethod.DYNAMIC)
                    .withSubnet(new SubnetInner().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet"))))
                .withBackendAddressPools(Arrays.asList()).withLoadBalancingRules(Arrays.asList())
                .withProbes(Arrays.asList()).withInboundNatRules(Arrays.asList())
                .withInboundNatPools(Arrays.asList(new InboundNatPool().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test")
                    .withName("test")
                    .withFrontendIpConfiguration(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"))
                    .withProtocol(TransportProtocol.TCP).withFrontendPortRangeStart(8080).withFrontendPortRangeEnd(8085)
                    .withBackendPort(8888).withIdleTimeoutInMinutes(10).withEnableFloatingIp(true)
                    .withEnableTcpReset(true)))
                .withOutboundRules(Arrays.asList()),
            com.azure.core.util.Context.NONE);
    }
}
