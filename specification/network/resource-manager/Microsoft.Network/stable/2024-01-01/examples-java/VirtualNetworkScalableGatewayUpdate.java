
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayNatRuleInner;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySku;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySkuName;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySkuTier;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewayType;
import com.azure.resourcemanager.network.models.VpnNatRuleMapping;
import com.azure.resourcemanager.network.models.VpnNatRuleMode;
import com.azure.resourcemanager.network.models.VpnNatRuleType;
import com.azure.resourcemanager.network.models.VpnType;
import java.util.Arrays;

/**
 * Samples for VirtualNetworkGateways CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/
     * VirtualNetworkScalableGatewayUpdate.json
     */
    /**
     * Sample code: UpdateVirtualNetworkScalableGateway.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualNetworkScalableGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkGateways().createOrUpdate("rg1", "ergw",
            new VirtualNetworkGatewayInner().withLocation("centralus")
                .withIpConfigurations(Arrays.asList(new VirtualNetworkGatewayIpConfigurationInner()
                    .withName("gwipconfig1").withPrivateIpAllocationMethod(IpAllocationMethod.STATIC)
                    .withSubnet(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"))
                    .withPublicIpAddress(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"))))
                .withGatewayType(VirtualNetworkGatewayType.EXPRESS_ROUTE).withVpnType(VpnType.POLICY_BASED)
                .withEnableBgp(false).withActive(false).withDisableIpSecReplayProtection(false)
                .withSku(new VirtualNetworkGatewaySku().withName(VirtualNetworkGatewaySkuName.ER_GW_SCALE)
                    .withTier(VirtualNetworkGatewaySkuTier.ER_GW_SCALE))
                .withNatRules(Arrays.asList(new VirtualNetworkGatewayNatRuleInner().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule1")
                    .withName("natRule1").withTypePropertiesType(VpnNatRuleType.STATIC)
                    .withMode(VpnNatRuleMode.EGRESS_SNAT)
                    .withInternalMappings(Arrays.asList(new VpnNatRuleMapping().withAddressSpace("10.10.0.0/24")))
                    .withExternalMappings(Arrays.asList(new VpnNatRuleMapping().withAddressSpace("50.0.0.0/24")))
                    .withIpConfigurationId(""),
                    new VirtualNetworkGatewayNatRuleInner().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/ergw/natRules/natRule2")
                        .withName("natRule2").withTypePropertiesType(VpnNatRuleType.STATIC)
                        .withMode(VpnNatRuleMode.INGRESS_SNAT)
                        .withInternalMappings(Arrays.asList(new VpnNatRuleMapping().withAddressSpace("20.10.0.0/24")))
                        .withExternalMappings(Arrays.asList(new VpnNatRuleMapping().withAddressSpace("30.0.0.0/24")))
                        .withIpConfigurationId("")))
                .withEnableBgpRouteTranslationForNat(false).withAllowVirtualWanTraffic(false)
                .withAllowRemoteVnetTraffic(false),
            com.azure.core.util.Context.NONE);
    }
}
