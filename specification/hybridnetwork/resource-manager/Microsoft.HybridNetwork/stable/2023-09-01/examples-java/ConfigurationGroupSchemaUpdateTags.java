
import com.azure.resourcemanager.hybridnetwork.models.ConfigurationGroupSchema;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ConfigurationGroupSchemas Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupSchemaUpdateTags.json
     */
    /**
     * Sample code: Create or update the configuration group schema resource.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateTheConfigurationGroupSchemaResource(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        ConfigurationGroupSchema resource = manager.configurationGroupSchemas()
            .getWithResponse("rg1", "testPublisher", "testConfigurationGroupSchema", com.azure.core.util.Context.NONE)
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
