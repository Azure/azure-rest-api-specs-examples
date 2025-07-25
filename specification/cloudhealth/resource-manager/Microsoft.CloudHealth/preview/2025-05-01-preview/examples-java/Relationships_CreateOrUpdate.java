
import com.azure.resourcemanager.cloudhealth.models.RelationshipProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Relationships CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Relationships_CreateOrUpdate.json
     */
    /**
     * Sample code: Relationships_CreateOrUpdate.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void relationshipsCreateOrUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.relationships().define("rel1").withExistingHealthmodel("rgopenapi", "model1")
            .withProperties(
                new RelationshipProperties().withDisplayName("My relationship").withParentEntityName("Entity1")
                    .withChildEntityName("Entity2").withLabels(mapOf("key9681", "fakeTokenPlaceholder")))
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
