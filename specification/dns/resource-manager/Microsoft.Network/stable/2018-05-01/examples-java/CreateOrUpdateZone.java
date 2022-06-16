import com.azure.core.util.Context;
import com.azure.resourcemanager.dns.fluent.models.ZoneInner;
import java.util.HashMap;
import java.util.Map;

/** Samples for Zones CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateZone.json
     */
    /**
     * Sample code: Create zone.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createZone(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .dnsZones()
            .manager()
            .serviceClient()
            .getZones()
            .createOrUpdateWithResponse(
                "rg1",
                "zone1",
                new ZoneInner().withLocation("Global").withTags(mapOf("key1", "value1")),
                null,
                null,
                Context.NONE);
    }

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
