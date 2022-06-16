import com.azure.core.management.SubResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for DnsResolvers CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsResolver_Put.json
     */
    /**
     * Sample code: Upsert DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void upsertDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .dnsResolvers()
            .define("sampleDnsResolver")
            .withRegion("westus2")
            .withExistingResourceGroup("sampleResourceGroup")
            .withTags(mapOf("key1", "value1"))
            .withVirtualNetwork(
                new SubResource()
                    .withId(
                        "/subscriptions/cbb1387e-4b03-44f2-ad41-58d4677b9873/resourceGroups/virtualNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork"))
            .create();
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
