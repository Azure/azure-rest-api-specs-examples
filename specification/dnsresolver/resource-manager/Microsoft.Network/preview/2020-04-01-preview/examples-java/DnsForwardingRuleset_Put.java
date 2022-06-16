import com.azure.core.management.SubResource;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for DnsForwardingRulesets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/DnsForwardingRuleset_Put.json
     */
    /**
     * Sample code: Upsert DNS forwarding ruleset.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void upsertDNSForwardingRuleset(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .dnsForwardingRulesets()
            .define("samplednsForwardingRuleset")
            .withRegion("westus2")
            .withExistingResourceGroup("sampleResourceGroup")
            .withTags(mapOf("key1", "value1"))
            .withDnsResolverOutboundEndpoints(
                Arrays
                    .asList(
                        new SubResource()
                            .withId(
                                "/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint0"),
                        new SubResource()
                            .withId(
                                "/subscriptions/abdd4249-9f34-4cc6-8e42-c2e32110603e/resourceGroups/sampleResourceGroup/providers/Microsoft.Network/dnsResolvers/sampleDnsResolver/outboundEndpoints/sampleOutboundEndpoint1")))
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
