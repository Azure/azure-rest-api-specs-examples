
import com.azure.resourcemanager.oracledatabase.models.ResourceAnchorProperties;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ResourceAnchors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ResourceAnchors_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: ResourceAnchors_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void resourceAnchorsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.resourceAnchors().define("resourceanchor1").withRegion("at").withExistingResourceGroup("rgopenapi")
            .withTags(mapOf("key236", "fakeTokenPlaceholder")).withProperties(new ResourceAnchorProperties()).create();
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
