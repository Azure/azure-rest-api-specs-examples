
import com.azure.resourcemanager.networksecurity.models.NeuroShieldPatchProperties;
import com.azure.resourcemanager.networksecurity.models.NeuroShieldPatchTags;
import com.azure.resourcemanager.networksecurity.models.NeuroShieldResource;
import com.azure.resourcemanager.networksecurity.models.NeuroShieldResponseMode;
import com.azure.resourcemanager.networksecurity.models.NeuroShieldSeverityThreshold;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for NeuroShields Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/NeuroShields_Update.json
     */
    /**
     * Sample code: Update NeuroShield resource.
     * 
     * @param manager Entry point to NetworkSecurityManager.
     */
    public static void
        updateNeuroShieldResource(com.azure.resourcemanager.networksecurity.NetworkSecurityManager manager) {
        NeuroShieldResource resource = manager.neuroShields()
            .getByResourceGroupWithResponse("myResourceGroup", "myNeuroShield", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(new NeuroShieldPatchTags().withAdditionalProperties(mapOf("env", "production")))
            .withProperties(new NeuroShieldPatchProperties().withResponseMode(NeuroShieldResponseMode.AUTO)
                .withSeverityThreshold(NeuroShieldSeverityThreshold.HIGH)
                .withDescription("Updated NeuroShield resource"))
            .withIfMatch("\"neuroshield-etag\"").apply();
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
