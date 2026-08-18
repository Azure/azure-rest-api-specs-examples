
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.BackendAddressPoolInner;
import com.azure.resourcemanager.network.fluent.models.FrontendIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancerInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancingRuleInner;
import com.azure.resourcemanager.network.fluent.models.ProbeInner;
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.models.LoadBalancerMode;
import com.azure.resourcemanager.network.models.LoadBalancerScope;
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
     * x-ms-original-file: 2025-09-01/LoadBalancerCreateWithAdvancedMode.json
     */
    /**
     * Sample code: Create load balancer with advanced mode.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createLoadBalancerWithAdvancedMode(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getLoadBalancers().createOrUpdate("rg1", "lb", new LoadBalancerInner()
            .withLocation("eastus").withSku(new LoadBalancerSku().withName(LoadBalancerSkuName.STANDARD))
            .withFrontendIpConfigurations(Arrays.asList(new FrontendIpConfigurationInner().withName("fe-lb")
                .withSubnet(new SubnetInner().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"))
                .withEnableConnectionTracking(true)))
            .withBackendAddressPools(Arrays.asList(new BackendAddressPoolInner().withName("be-lb")))
            .withLoadBalancingRules(Arrays.asList(new LoadBalancingRuleInner().withName("rulelb")
                .withFrontendIpConfiguration(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"))
                .withBackendAddressPool(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"))
                .withProbe(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"))
                .withProtocol(TransportProtocol.UDP).withLoadDistribution(LoadDistribution.DEFAULT)
                .withFrontendPort(4789).withBackendPort(4789).withEnableFloatingIp(true)))
            .withProbes(Arrays.asList(new ProbeInner().withName("probe-lb").withProtocol(ProbeProtocol.HTTP)
                .withPort(80).withIntervalInSeconds(15).withNumberOfProbes(2).withProbeThreshold(1)
                .withRequestPath("healthcheck.aspx")))
            .withScope(LoadBalancerScope.PUBLIC).withMode(LoadBalancerMode.ADVANCED), com.azure.core.util.Context.NONE);
    }
}
