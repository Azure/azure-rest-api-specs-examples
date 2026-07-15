
import com.azure.resourcemanager.cloudhealth.models.HealthModelProperties;
import com.azure.resourcemanager.cloudhealth.models.ManagedServiceIdentity;
import com.azure.resourcemanager.cloudhealth.models.ManagedServiceIdentityType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for HealthModels Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/HealthModels_Create.json
     */
    /**
     * Sample code: HealthModels_Create.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void healthModelsCreate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().define("online-store").withRegion("eastus").withExistingResourceGroup("online-store-rg")
            .withTags(mapOf("environment", "production", "team", "online-store"))
            .withProperties(new HealthModelProperties())
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)).create();
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
