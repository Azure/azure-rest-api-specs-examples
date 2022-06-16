import com.azure.core.management.SubResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for OutboundEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/OutboundEndpoint_Put.json
     */
    /**
     * Sample code: Upsert outbound endpoint for DNS resolver.
     *
     * @param manager Entry point to DnsResolverManager.
     */
    public static void upsertOutboundEndpointForDNSResolver(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager
            .outboundEndpoints()
            .define("sampleOutboundEndpoint")
            .withRegion("westus2")
            .withExistingDnsResolver("sampleResourceGroup", "sampleDnsResolver")
            .withTags(mapOf("key1", "value1"))
            .withSubnet(
                new SubResource()
                    .withId(
                        "/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"))
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
