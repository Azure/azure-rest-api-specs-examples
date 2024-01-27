
import com.azure.resourcemanager.search.models.IpRule;
import com.azure.resourcemanager.search.models.NetworkRuleSet;
import com.azure.resourcemanager.search.models.PublicNetworkAccess;
import com.azure.resourcemanager.search.models.SearchServiceUpdate;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Services Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/search/resource-manager/Microsoft.Search/stable/2023-11-01/examples/
     * SearchUpdateServiceToAllowAccessFromPublicCustomIPs.json
     */
    /**
     * Sample code: SearchUpdateServiceToAllowAccessFromPublicCustomIPs.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        searchUpdateServiceToAllowAccessFromPublicCustomIPs(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.searchServices().manager().serviceClient().getServices().updateWithResponse("rg1", "mysearchservice",
            new SearchServiceUpdate().withReplicaCount(3).withPartitionCount(1)
                .withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
                .withNetworkRuleSet(new NetworkRuleSet().withIpRules(
                    Arrays.asList(new IpRule().withValue("123.4.5.6"), new IpRule().withValue("123.4.6.0/18")))),
            null, com.azure.core.util.Context.NONE);
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
