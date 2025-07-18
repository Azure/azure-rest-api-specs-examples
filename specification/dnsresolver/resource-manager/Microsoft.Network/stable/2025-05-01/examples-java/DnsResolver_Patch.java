
import com.azure.resourcemanager.dnsresolver.models.DnsResolver;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DnsResolvers Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/DnsResolver_Patch.json
     */
    /**
     * Sample code: Update DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void updateDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        DnsResolver resource = manager.dnsResolvers().getByResourceGroupWithResponse("sampleResourceGroup",
            "sampleDnsResolver", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key1", "fakeTokenPlaceholder")).apply();
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
