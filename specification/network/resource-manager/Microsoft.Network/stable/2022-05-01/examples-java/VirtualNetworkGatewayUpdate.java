import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayNatRuleInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import com.azure.resourcemanager.network.models.BgpSettings;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.RadiusServer;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySku;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySkuName;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySkuTier;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewayType;
import com.azure.resourcemanager.network.models.VpnClientConfiguration;
import com.azure.resourcemanager.network.models.VpnClientProtocol;
import com.azure.resourcemanager.network.models.VpnNatRuleMapping;
import com.azure.resourcemanager.network.models.VpnNatRuleMode;
import com.azure.resourcemanager.network.models.VpnNatRuleType;
import com.azure.resourcemanager.network.models.VpnType;
import java.util.Arrays;

/** Samples for VirtualNetworkGateways CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/VirtualNetworkGatewayUpdate.json
     */
    /**
     * Sample code: UpdateVirtualNetworkGateway.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateVirtualNetworkGateway(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getVirtualNetworkGateways()
            .createOrUpdate(
                "rg1",
                "vpngw",
                new VirtualNetworkGatewayInner()
                    .withLocation("centralus")
                    .withIpConfigurations(
                        Arrays
                            .asList(
                                new VirtualNetworkGatewayIpConfigurationInner()
                                    .withName("gwipconfig1")
                                    .withPrivateIpAllocationMethod(IpAllocationMethod.DYNAMIC)
                                    .withSubnet(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"))
                                    .withPublicIpAddress(
                                        new SubResource()
                                            .withId(
                                                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"))))
                    .withGatewayType(VirtualNetworkGatewayType.VPN)
                    .withVpnType(VpnType.ROUTE_BASED)
                    .withEnableBgp(false)
                    .withActive(false)
                    .withDisableIpSecReplayProtection(false)
                    .withSku(
                        new VirtualNetworkGatewaySku()
                            .withName(VirtualNetworkGatewaySkuName.VPN_GW1)
                            .withTier(VirtualNetworkGatewaySkuTier.VPN_GW1))
                    .withVpnClientConfiguration(
                        new VpnClientConfiguration()
                            .withVpnClientRootCertificates(Arrays.asList())
                            .withVpnClientRevokedCertificates(Arrays.asList())
                            .withVpnClientProtocols(Arrays.asList(VpnClientProtocol.OPEN_VPN))
                            .withRadiusServers(
                                Arrays
                                    .asList(
                                        new RadiusServer()
                                            .withRadiusServerAddress("10.2.0.0")
                                            .withRadiusServerScore(20L)
                                            .withRadiusServerSecret("radiusServerSecret"))))
                    .withBgpSettings(
                        new BgpSettings().withAsn(65515L).withBgpPeeringAddress("10.0.1.30").withPeerWeight(0))
                    .withCustomRoutes(new AddressSpace().withAddressPrefixes(Arrays.asList("101.168.0.6/32")))
                    .withEnableDnsForwarding(true)
                    .withNatRules(
                        Arrays
                            .asList(
                                new VirtualNetworkGatewayNatRuleInner()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1")
                                    .withName("natRule1")
                                    .withTypePropertiesType(VpnNatRuleType.STATIC)
                                    .withMode(VpnNatRuleMode.EGRESS_SNAT)
                                    .withInternalMappings(
                                        Arrays.asList(new VpnNatRuleMapping().withAddressSpace("10.10.0.0/24")))
                                    .withExternalMappings(
                                        Arrays.asList(new VpnNatRuleMapping().withAddressSpace("50.0.0.0/24")))
                                    .withIpConfigurationId(""),
                                new VirtualNetworkGatewayNatRuleInner()
                                    .withId(
                                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2")
                                    .withName("natRule2")
                                    .withTypePropertiesType(VpnNatRuleType.STATIC)
                                    .withMode(VpnNatRuleMode.INGRESS_SNAT)
                                    .withInternalMappings(
                                        Arrays.asList(new VpnNatRuleMapping().withAddressSpace("20.10.0.0/24")))
                                    .withExternalMappings(
                                        Arrays.asList(new VpnNatRuleMapping().withAddressSpace("30.0.0.0/24")))
                                    .withIpConfigurationId("")))
                    .withEnableBgpRouteTranslationForNat(false)
                    .withAllowVirtualWanTraffic(false)
                    .withAllowRemoteVnetTraffic(false),
                Context.NONE);
    }
}
