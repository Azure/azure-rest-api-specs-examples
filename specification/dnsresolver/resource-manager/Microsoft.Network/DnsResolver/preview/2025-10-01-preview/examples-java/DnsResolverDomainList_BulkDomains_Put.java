
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DnsResolverDomainLists CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/DnsResolverDomainList_BulkDomains_Put.json
     */
    /**
     * Sample code: Upsert DNS resolver domain list with bulk number of domains.
     * 
     * @param manager Entry point to DnsResolverManager.
     */
    public static void upsertDNSResolverDomainListWithBulkNumberOfDomains(
        com.azure.resourcemanager.dnsresolver.DnsResolverManager manager) {
        manager.dnsResolverDomainLists().define("sampleDnsResolverDomainList").withRegion("westus2")
            .withExistingResourceGroup("sampleResourceGroup").withTags(mapOf("key1", "fakeTokenPlaceholder")).create();
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
