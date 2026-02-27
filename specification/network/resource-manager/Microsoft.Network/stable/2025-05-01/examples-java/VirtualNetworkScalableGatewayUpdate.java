
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkGatewayIpConfigurationInner;
import com.azure.resourcemanager.network.models.AdminState;
import com.azure.resourcemanager.network.models.IpAllocationMethod;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewayAutoScaleBounds;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewayAutoScaleConfiguration;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySku;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySkuName;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewaySkuTier;
import com.azure.resourcemanager.network.models.VirtualNetworkGatewayType;
import com.azure.resourcemanager.network.models.VpnType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualNetworkGateways CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
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
                .withAutoScaleConfiguration(new VirtualNetworkGatewayAutoScaleConfiguration()
                    .withBounds(new VirtualNetworkGatewayAutoScaleBounds().withMin(2).withMax(3)))
                .withIpConfigurations(Arrays.asList(new VirtualNetworkGatewayIpConfigurationInner()
                    .withName("gwipconfig1").withPrivateIpAllocationMethod(IpAllocationMethod.DYNAMIC)
                    .withSubnet(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"))
                    .withPublicIpAddress(new SubResource().withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"))))
                .withGatewayType(VirtualNetworkGatewayType.EXPRESS_ROUTE).withVpnType(VpnType.POLICY_BASED)
                .withEnableBgp(false).withActive(false).withDisableIpSecReplayProtection(false)
                .withSku(new VirtualNetworkGatewaySku().withName(VirtualNetworkGatewaySkuName.ER_GW_SCALE)
                    .withTier(VirtualNetworkGatewaySkuTier.ER_GW_SCALE))
                .withVirtualNetworkGatewayPolicyGroups(Arrays.asList()).withNatRules(Arrays.asList())
                .withEnableBgpRouteTranslationForNat(false).withAllowVirtualWanTraffic(false)
                .withAllowRemoteVnetTraffic(false).withAdminState(AdminState.ENABLED),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
