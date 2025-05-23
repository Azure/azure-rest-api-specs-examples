
import com.azure.resourcemanager.hybridnetwork.models.NetworkFunctionDefinitionGroup;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NetworkFunctionDefinitionGroups Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * NetworkFunctionDefinitionGroupUpdateTags.json
     */
    /**
     * Sample code: Create or update the network function definition group resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateTheNetworkFunctionDefinitionGroupResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        NetworkFunctionDefinitionGroup resource = manager.networkFunctionDefinitionGroups().getWithResponse("rg",
            "TestPublisher", "TestNetworkFunctionDefinitionGroupName", com.azure.core.util.Context.NONE).getValue();
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
