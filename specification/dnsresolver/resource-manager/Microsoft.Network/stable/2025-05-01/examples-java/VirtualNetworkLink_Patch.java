
import com.azure.resourcemanager.dnsresolver.models.VirtualNetworkLink;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for VirtualNetworkLinks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dnsresolver/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualNetworkLink_Patch.
     * json
     */
    /**
     * Sample code: Update virtual network link to a DNS forwarding ruleset.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void updateVirtualNetworkLinkToADNSForwardingRuleset(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        VirtualNetworkLink resource = manager.virtualNetworkLinks().getWithResponse("sampleResourceGroup",
            "sampleDnsForwardingRuleset", "sampleVirtualNetworkLink", com.azure.core.util.Context.NONE).getValue();
        resource.update().withMetadata(mapOf("additionalProp1", "value1")).apply();
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
