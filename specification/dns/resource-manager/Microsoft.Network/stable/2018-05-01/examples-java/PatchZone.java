
import com.azure.resourcemanager.dns.models.ZoneUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Zones Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/PatchZone.json
     */
    /**
     * Sample code: Patch zone.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.dnsZones().manager().serviceClient().getZones().updateWithResponse("rg1", "zone1",
            new ZoneUpdate().withTags(mapOf("key2", "fakeTokenPlaceholder")), null, com.azure.core.util.Context.NONE);
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
