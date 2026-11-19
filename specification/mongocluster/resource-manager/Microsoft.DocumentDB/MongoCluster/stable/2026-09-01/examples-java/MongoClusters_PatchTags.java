
import com.azure.resourcemanager.mongocluster.models.MongoCluster;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for MongoClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-01/MongoClusters_PatchTags.json
     */
    /**
     * Sample code: Updates string-valued resource tags without changing the existing PremiumSSD storage type.
     * 
     * @param manager Entry point to MongoClusterManager.
     */
    public static void updatesStringValuedResourceTagsWithoutChangingTheExistingPremiumSSDStorageType(
        com.azure.resourcemanager.mongocluster.MongoClusterManager manager) {
        MongoCluster resource = manager.mongoClusters()
            .getByResourceGroupWithResponse("TestResourceGroup", "myMongoCluster", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("environment", "production", "costCenter", "1234")).apply();
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
