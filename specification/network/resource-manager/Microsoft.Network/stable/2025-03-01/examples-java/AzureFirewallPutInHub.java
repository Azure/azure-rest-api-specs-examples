
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.network.fluent.models.AzureFirewallInner;
import com.azure.resourcemanager.network.models.AzureFirewallSku;
import com.azure.resourcemanager.network.models.AzureFirewallSkuName;
import com.azure.resourcemanager.network.models.AzureFirewallSkuTier;
import com.azure.resourcemanager.network.models.AzureFirewallThreatIntelMode;
import com.azure.resourcemanager.network.models.HubIpAddresses;
import com.azure.resourcemanager.network.models.HubPublicIpAddresses;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for AzureFirewalls CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/AzureFirewallPutInHub.json
     */
    /**
     * Sample code: Create Azure Firewall in virtual Hub.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAzureFirewallInVirtualHub(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getAzureFirewalls().createOrUpdate("rg1", "azurefirewall",
            new AzureFirewallInner().withLocation("West US").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withZones(Arrays.asList()).withThreatIntelMode(AzureFirewallThreatIntelMode.ALERT)
                .withVirtualHub(new SubResource()
                    .withId("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"))
                .withFirewallPolicy(new SubResource().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1"))
                .withHubIpAddresses(new HubIpAddresses()
                    .withPublicIPs(new HubPublicIpAddresses().withAddresses(Arrays.asList()).withCount(1)))
                .withSku(new AzureFirewallSku().withName(AzureFirewallSkuName.AZFW_HUB)
                    .withTier(AzureFirewallSkuTier.STANDARD)),
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
