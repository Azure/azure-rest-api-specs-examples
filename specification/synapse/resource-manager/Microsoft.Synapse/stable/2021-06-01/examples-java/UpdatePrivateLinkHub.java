
import com.azure.resourcemanager.synapse.models.PrivateLinkHub;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateLinkHubs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/UpdatePrivateLinkHub.json
     */
    /**
     * Sample code: Update a privateLinkHub.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void updateAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        PrivateLinkHub resource = manager.privateLinkHubs()
            .getByResourceGroupWithResponse("resourceGroup1", "privateLinkHub1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key", "fakeTokenPlaceholder")).apply();
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
