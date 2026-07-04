
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.AzureFirewallInner;
import com.azure.resourcemanager.network.models.AzureFirewallIpConfiguration;
import com.azure.resourcemanager.network.models.AzureFirewallSku;
import com.azure.resourcemanager.network.models.AzureFirewallSkuName;
import com.azure.resourcemanager.network.models.AzureFirewallSkuTier;
import com.azure.resourcemanager.network.models.AzureFirewallThreatIntelMode;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AzureFirewalls CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AzureFirewallPutWithAfcConfiguration.json
     */
    /**
     * Sample code: Create Azure Firewall With AFC Control Plane.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createAzureFirewallWithAFCControlPlane(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getAzureFirewalls().createOrUpdate("rg1", "azurefirewall", new AzureFirewallInner()
            .withLocation("West US").withTags(mapOf("key1", "fakeTokenPlaceholder")).withZones(Arrays.asList())
            .withIpConfigurations(Arrays.asList(new AzureFirewallIpConfiguration()
                .withName("azureFirewallIpConfiguration")
                .withSubnet(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"))
                .withPublicIpAddress(new SubResource().withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"))))
            .withThreatIntelMode(AzureFirewallThreatIntelMode.ALERT).withSku(new AzureFirewallSku()
                .withName(AzureFirewallSkuName.AZFW_VNET).withTier(AzureFirewallSkuTier.STANDARD)),
            true, com.azure.core.util.Context.NONE);
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
