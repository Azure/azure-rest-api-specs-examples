
import com.azure.resourcemanager.privatedns.fluent.models.PrivateZoneInner;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateZones Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZonePatch.json
     */
    /**
     * Sample code: PATCH Private DNS Zone.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void pATCHPrivateDNSZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.privateDnsZones().manager().serviceClient().getPrivateZones().update("resourceGroup1", "privatezone1.com",
            new PrivateZoneInner().withTags(mapOf("key2", "fakeTokenPlaceholder")), null,
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
