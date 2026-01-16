
import com.azure.resourcemanager.resourceconnector.models.Appliance;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Appliances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/AppliancesPatch.json
     */
    /**
     * Sample code: Update Appliance.
     * 
     * @param manager Entry point to ResourceConnectorManager.
     */
    public static void updateAppliance(com.azure.resourcemanager.resourceconnector.ResourceConnectorManager manager) {
        Appliance resource = manager.appliances()
            .getByResourceGroupWithResponse("testresourcegroup", "appliance01", com.azure.core.util.Context.NONE)
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
