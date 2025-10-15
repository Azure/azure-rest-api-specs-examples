
import com.azure.resourcemanager.oracledatabase.models.CloudExadataInfrastructureProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CloudExadataInfrastructures CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/exaInfra_create.json
     */
    /**
     * Sample code: CloudExadataInfrastructures_createOrUpdate.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudExadataInfrastructuresCreateOrUpdate(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().define("infra1").withRegion("eastus").withExistingResourceGroup("rg000")
            .withZones(Arrays.asList("1")).withTags(mapOf("tagK1", "tagV1"))
            .withProperties(new CloudExadataInfrastructureProperties().withComputeCount(100).withStorageCount(10)
                .withShape("EXADATA.X9M").withDisplayName("infra 1"))
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
