
import com.azure.resourcemanager.privatedns.fluent.models.PrivateZoneInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateZones CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/PrivateZonePut.json
     */
    /**
     * Sample code: PUT Private DNS Zone.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pUTPrivateDNSZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getPrivateZones().createOrUpdate("resourceGroup1",
            "privatezone1.com",
            new PrivateZoneInner().withLocation("Global").withTags(mapOf("key1", "fakeTokenPlaceholder")), null, null,
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
