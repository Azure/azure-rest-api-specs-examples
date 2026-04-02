
import com.azure.resourcemanager.search.models.IpRule;
import com.azure.resourcemanager.search.models.NetworkRuleSet;
import com.azure.resourcemanager.search.models.PublicNetworkAccess;
import com.azure.resourcemanager.search.models.SearchBypass;
import com.azure.resourcemanager.search.models.SearchServiceUpdate;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Services Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/SearchUpdateServiceToAllowAccessFromPublicCustomIPs.json
     */
    /**
     * Sample code: SearchUpdateServiceToAllowAccessFromPublicCustomIPs.
     * 
     * @param manager Entry point to SearchServiceManager.
     */
    public static void searchUpdateServiceToAllowAccessFromPublicCustomIPs(
        com.azure.resourcemanager.search.SearchServiceManager manager) {
        manager.serviceClient().getServices().updateWithResponse("rg1", "mysearchservice", new SearchServiceUpdate()
            .withReplicaCount(3).withPartitionCount(1).withPublicNetworkAccess(PublicNetworkAccess.ENABLED)
            .withNetworkRuleSet(new NetworkRuleSet()
                .withIpRules(Arrays.asList(new IpRule().withValue("123.4.5.6"), new IpRule().withValue("123.4.6.0/18")))
                .withBypass(SearchBypass.NONE)),
            com.azure.core.util.Context.NONE);
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
