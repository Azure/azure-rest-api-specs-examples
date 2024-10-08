
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.privatedns.fluent.models.VirtualNetworkLinkInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualNetworkLinks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/VirtualNetworkLinkPut.json
     */
    /**
     * Sample code: PUT Private DNS Zone Virtual Network Link.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pUTPrivateDNSZoneVirtualNetworkLink(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getVirtualNetworkLinks().createOrUpdate("resourceGroup1",
            "privatezone1.com", "virtualNetworkLink1",
            new VirtualNetworkLinkInner().withLocation("Global").withTags(mapOf("key1", "fakeTokenPlaceholder"))
                .withVirtualNetwork(new SubResource().withId(
                    "/subscriptions/virtualNetworkSubscriptionId/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/virtualNetworkName"))
                .withRegistrationEnabled(false),
            null, null, com.azure.core.util.Context.NONE);
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
