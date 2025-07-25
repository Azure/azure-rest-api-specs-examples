
import com.azure.resourcemanager.cloudhealth.models.HealthModel;
import com.azure.resourcemanager.cloudhealth.models.ManagedServiceIdentity;
import com.azure.resourcemanager.cloudhealth.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.cloudhealth.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for HealthModels Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/HealthModels_Update.json
     */
    /**
     * Sample code: HealthModels_Update.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void healthModelsUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        HealthModel resource = manager.healthModels()
            .getByResourceGroupWithResponse("rgopenapi", "model1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withTags(mapOf("key21", "fakeTokenPlaceholder")).withIdentity(new ManagedServiceIdentity()
            .withType(ManagedServiceIdentityType.fromString("SystemAssigned, UserAssigned"))
            .withUserAssignedIdentities(mapOf(
                "/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ua1",
                new UserAssignedIdentity())))
            .apply();
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
