
import com.azure.resourcemanager.cloudhealth.models.DiscoveryRuleRecommendedSignalsBehavior;
import com.azure.resourcemanager.cloudhealth.models.HealthModelProperties;
import com.azure.resourcemanager.cloudhealth.models.ManagedServiceIdentity;
import com.azure.resourcemanager.cloudhealth.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.cloudhealth.models.ModelDiscoverySettings;
import com.azure.resourcemanager.cloudhealth.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for HealthModels Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/HealthModels_Create.json
     */
    /**
     * Sample code: HealthModels_Create.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void healthModelsCreate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.healthModels().define("model1").withRegion("eastus2").withExistingResourceGroup("rgopenapi")
            .withTags(mapOf("key2961", "fakeTokenPlaceholder"))
            .withProperties(new HealthModelProperties().withDiscovery(
                new ModelDiscoverySettings().withScope("/providers/Microsoft.Management/serviceGroups/myServiceGroup")
                    .withAddRecommendedSignals(DiscoveryRuleRecommendedSignalsBehavior.ENABLED)
                    .withIdentity("SystemAssigned")))
            .withIdentity(new ManagedServiceIdentity()
                .withType(ManagedServiceIdentityType.fromString("SystemAssigned, UserAssigned"))
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/4980D7D5-4E07-47AD-AD34-E76C6BC9F061/resourceGroups/rgopenapi/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ua1",
                    new UserAssignedIdentity())))
            .create();
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
