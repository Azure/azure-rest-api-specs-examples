Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.13.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.BackendAddressPoolInner;
import com.azure.resourcemanager.network.fluent.models.FrontendIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.InboundNatRuleInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancerInner;
import com.azure.resourcemanager.network.fluent.models.LoadBalancingRuleInner;
import com.azure.resourcemanager.network.fluent.models.OutboundRuleInner;
import com.azure.resourcemanager.network.fluent.models.ProbeInner;
import com.azure.resourcemanager.network.fluent.models.PublicIpAddressInner;
import com.azure.resourcemanager.network.models.LoadBalancerOutboundRuleProtocol;
import com.azure.resourcemanager.network.models.LoadBalancerSku;
import com.azure.resourcemanager.network.models.LoadBalancerSkuName;
import com.azure.resourcemanager.network.models.LoadDistribution;
import com.azure.resourcemanager.network.models.ProbeProtocol;
import com.azure.resourcemanager.network.models.TransportProtocol;
import java.util.Arrays;

/** Samples for LoadBalancers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerCreateWithOutboundRules.json
     */
    /**
     * Sample code: Create load balancer with outbound rules.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createLoadBalancerWithOutboundRules(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getLoadBalancers()
            .createOrUpdate(
                "rg1",
                "lb",
                new LoadBalancerInner()
                    .withLocation("eastus")
                    .withSku(new LoadBalancerSku().withName(LoadBalancerSkuName.STANDARD))
                    .withFrontendIpConfigurations(
                        Arrays
                            .asList(
                                new FrontendIpConfigurationInner()
                                    .withName("fe-lb")
                                    .withPublicIpAddress(
                                        new PublicIpAddressInner()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pip"))))
                    .withBackendAddressPools(Arrays.asList(new BackendAddressPoolInner().withName("be-lb")))
                    .withLoadBalancingRules(
                        Arrays
                            .asList(
                                new LoadBalancingRuleInner()
                                    .withName("rulelb")
                                    .withFrontendIpConfiguration(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"))
                                    .withBackendAddressPool(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"))
                                    .withProbe(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"))
                                    .withProtocol(TransportProtocol.TCP)
                                    .withLoadDistribution(LoadDistribution.DEFAULT)
                                    .withFrontendPort(80)
                                    .withBackendPort(80)
                                    .withIdleTimeoutInMinutes(15)
                                    .withEnableFloatingIp(true)
                                    .withDisableOutboundSnat(true)))
                    .withProbes(
                        Arrays
                            .asList(
                                new ProbeInner()
                                    .withName("probe-lb")
                                    .withProtocol(ProbeProtocol.HTTP)
                                    .withPort(80)
                                    .withIntervalInSeconds(15)
                                    .withNumberOfProbes(2)
                                    .withRequestPath("healthcheck.aspx")))
                    .withInboundNatRules(
                        Arrays
                            .asList(
                                new InboundNatRuleInner()
                                    .withName("in-nat-rule")
                                    .withFrontendIpConfiguration(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"))
                                    .withProtocol(TransportProtocol.TCP)
                                    .withFrontendPort(3389)
                                    .withBackendPort(3389)
                                    .withIdleTimeoutInMinutes(15)
                                    .withEnableFloatingIp(true)))
                    .withInboundNatPools(Arrays.asList())
                    .withOutboundRules(
                        Arrays
                            .asList(
                                new OutboundRuleInner()
                                    .withName("rule1")
                                    .withFrontendIpConfigurations(
                                        Arrays
                                            .asList(
                                                new SubResource()
                                                    .withId(
                                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb")))
                                    .withBackendAddressPool(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"))
                                    .withProtocol(LoadBalancerOutboundRuleProtocol.ALL))),
                Context.NONE);
    }
}
```
