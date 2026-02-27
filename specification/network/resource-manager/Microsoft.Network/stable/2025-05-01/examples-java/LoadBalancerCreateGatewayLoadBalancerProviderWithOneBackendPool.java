
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BackendAddressPoolInner;
import com.azure.resourcemanager.network.fluent.models.FrontendIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancerInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancingRuleInner;
import com.azure.resourcemanager.network.fluent.models.ProbeInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.GatewayLoadBalancerTunnelInterface;
import com.azure.resourcemanager.network.models.GatewayLoadBalancerTunnelInterfaceType;
import com.azure.resourcemanager.network.models.GatewayLoadBalancerTunnelProtocol;
import com.azure.resourcemanager.network.models.LoadBalancerSku;
import com.azure.resourcemanager.network.models.LoadBalancerSkuName;
import com.azure.resourcemanager.network.models.LoadDistribution;
import com.azure.resourcemanager.network.models.ProbeProtocol;
import com.azure.resourcemanager.network.models.TransportProtocol;
import java.util.Arrays;

/**
 * Samples for LoadBalancers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * LoadBalancerCreateGatewayLoadBalancerProviderWithOneBackendPool.json
     */
    /**
     * Sample code: Create load balancer with Gateway Load Balancer Provider configured with one Backend Pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createLoadBalancerWithGatewayLoadBalancerProviderConfiguredWithOneBackendPool(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().createOrUpdate("rg1", "lb",
            new LoadBalancerInner().withLocation("eastus")
                .withSku(new LoadBalancerSku().withName(LoadBalancerSkuName.GATEWAY))
                .withFrontendIpConfigurations(Arrays.asList(new FrontendIpConfigurationInner().withName("fe-lb")
                    .withSubnet(new SubnetInner().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"))))
                .withBackendAddressPools(Arrays.asList(new BackendAddressPoolInner().withName("be-lb")
                    .withTunnelInterfaces(Arrays.asList(
                        new GatewayLoadBalancerTunnelInterface().withPort(15000).withIdentifier(900)
                            .withProtocol(GatewayLoadBalancerTunnelProtocol.VXLAN)
                            .withType(GatewayLoadBalancerTunnelInterfaceType.INTERNAL),
                        new GatewayLoadBalancerTunnelInterface().withPort(15001).withIdentifier(901)
                            .withProtocol(GatewayLoadBalancerTunnelProtocol.VXLAN)
                            .withType(GatewayLoadBalancerTunnelInterfaceType.INTERNAL)))))
                .withLoadBalancingRules(Arrays.asList(new LoadBalancingRuleInner().withName("rulelb")
                    .withFrontendIpConfiguration(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"))
                    .withBackendAddressPools(Arrays.asList(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb")))
                    .withProbe(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"))
                    .withProtocol(TransportProtocol.ALL).withLoadDistribution(LoadDistribution.DEFAULT)
                    .withFrontendPort(0).withBackendPort(0).withIdleTimeoutInMinutes(15).withEnableFloatingIp(true)))
                .withProbes(Arrays.asList(new ProbeInner().withName("probe-lb").withProtocol(ProbeProtocol.HTTP)
                    .withPort(80).withIntervalInSeconds(15).withNumberOfProbes(2).withProbeThreshold(1)
                    .withRequestPath("healthcheck.aspx")))
                .withInboundNatPools(Arrays.asList()).withOutboundRules(Arrays.asList()),
            com.azure.core.util.Context.NONE);
    }
}
