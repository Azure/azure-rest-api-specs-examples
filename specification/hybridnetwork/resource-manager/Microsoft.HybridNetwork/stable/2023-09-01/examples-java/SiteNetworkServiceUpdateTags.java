
import com.azure.resourcemanager.hybridnetwork.models.SiteNetworkService;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for SiteNetworkServices UpdateTags.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * SiteNetworkServiceUpdateTags.json
     */
    /**
     * Sample code: Update network site tags.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void updateNetworkSiteTags(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        SiteNetworkService resource = manager.siteNetworkServices()
            .getByResourceGroupWithResponse("rg1", "testSiteNetworkServiceName", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
