
import com.azure.resourcemanager.servicelinker.models.ConfigurationInfo;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Connector GenerateConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/
     * GenerateConfigurations.json
     */
    /**
     * Sample code: GenerateConfiguration.
     * 
     * @param manager Entry point to ServiceLinkerManager.
     */
    public static void generateConfiguration(com.azure.resourcemanager.servicelinker.ServiceLinkerManager manager) {
        manager.connectors().generateConfigurationsWithResponse("00000000-0000-0000-0000-000000000000", "test-rg",
            "westus", "connectorName",
            new ConfigurationInfo().withCustomizedKeys(mapOf("ASL_DocumentDb_ConnectionString", "MyConnectionstring")),
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
