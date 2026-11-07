
import com.azure.resourcemanager.networksecurity.models.NeuroShieldProperties;
import com.azure.resourcemanager.networksecurity.models.NeuroShieldResponseMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NeuroShields CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShields_CreateOrUpdate_OmittedOptionalProperties.json
     */
    /**
     * Sample code: Create or update NeuroShield with omitted optional properties.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void createOrUpdateNeuroShieldWithOmittedOptionalProperties(
        com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        manager.neuroShields().define("myNeuroShield").withRegion("eastus2")
            .withExistingResourceGroup("myResourceGroup")
            .withProperties(new NeuroShieldProperties().withResponseMode(NeuroShieldResponseMode.MANUAL))
            .withIfMatch("\"neuroshield-etag\"").create();
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
