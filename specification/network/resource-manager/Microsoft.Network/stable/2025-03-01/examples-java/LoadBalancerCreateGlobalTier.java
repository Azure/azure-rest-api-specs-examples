
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BackendAddressPoolInner;
import com.azure.resourcemanager.network.fluent.models.FrontendIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancerInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancingRuleInner;
import com.azure.resourcemanager.network.fluent.models.ProbeInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.LoadBalancerBackendAddress;
import com.azure.resourcemanager.network.models.LoadBalancerSku;
import com.azure.resourcemanager.network.models.LoadBalancerSkuName;
import com.azure.resourcemanager.network.models.LoadBalancerSkuTier;
import com.azure.resourcemanager.network.models.LoadDistribution;
import com.azure.resourcemanager.network.models.ProbeProtocol;
import com.azure.resourcemanager.network.models.TransportProtocol;
import java.util.Arrays;

/**
 * Samples for LoadBalancers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerCreateGlobalTier.
     * json
     */
    /**
     * Sample code: Create load balancer with Global Tier and one regional load balancer in its backend pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createLoadBalancerWithGlobalTierAndOneRegionalLoadBalancerInItsBackendPool(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().createOrUpdate("rg1", "lb",
            new LoadBalancerInner().withLocation("eastus")
                .withSku(
                    new LoadBalancerSku().withName(LoadBalancerSkuName.STANDARD).withTier(LoadBalancerSkuTier.GLOBAL))
                .withFrontendIpConfigurations(Arrays.asList(new FrontendIpConfigurationInner().withName("fe-lb")
                    .withSubnet(new SubnetInner().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"))))
                .withBackendAddressPools(Arrays.asList(new BackendAddressPoolInner().withName("be-lb")
                    .withLoadBalancerBackendAddresses(Arrays.asList(new LoadBalancerBackendAddress()
                        .withName("regional-lb1-address")
                        .withLoadBalancerFrontendIpConfiguration(new SubResource().withId(
                            "/subscriptions/subid/resourceGroups/regional-lb-rg1/providers/Microsoft.Network/loadBalancers/regional-lb/frontendIPConfigurations/fe-rlb"))))))
                .withLoadBalancingRules(Arrays.asList(new LoadBalancingRuleInner().withName("rulelb")
                    .withFrontendIpConfiguration(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"))
                    .withBackendAddressPool(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"))
                    .withProbe(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"))
                    .withProtocol(TransportProtocol.TCP).withLoadDistribution(LoadDistribution.DEFAULT)
                    .withFrontendPort(80).withBackendPort(80).withIdleTimeoutInMinutes(15).withEnableFloatingIp(false)))
                .withProbes(Arrays.asList(new ProbeInner().withName("probe-lb").withProtocol(ProbeProtocol.HTTP)
                    .withPort(80).withIntervalInSeconds(15).withNumberOfProbes(2).withProbeThreshold(1)
                    .withRequestPath("healthcheck.aspx"))),
            com.azure.core.util.Context.NONE);
    }
}
