
import com.azure.resourcemanager.dnsresolver.models.OutboundEndpoint;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for OutboundEndpoints Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/OutboundEndpoint_Patch.
     * json
     */
    /**
     * Sample code: Update outbound endpoint for DNS resolver.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void
        updateOutboundEndpointForDNSResolver(com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        OutboundEndpoint resource = manager.outboundEndpoints().getWithResponse("sampleResourceGroup",
            "sampleDnsResolver", "sampleOutboundEndpoint", com.azure.core.util.Context.NONE).getValue();
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
