
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for PrivateLinkHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdatePrivateLinkHub.
     * json
     */
    /**
     * Sample code: Create or update a privateLinkHub.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateAPrivateLinkHub(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.privateLinkHubs().define("privateLinkHub1").withRegion("East US")
            .withExistingResourceGroup("resourceGroup1").withTags(mapOf("key", "fakeTokenPlaceholder")).create();
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
