
import com.azure.resourcemanager.oracledatabase.models.ExadbVmCluster;
import com.azure.resourcemanager.oracledatabase.models.ExadbVmClusterUpdateProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ExadbVmClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExadbVmClusters_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExadbVmClusters_Update_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exadbVmClustersUpdateMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        ExadbVmCluster resource = manager.exadbVmClusters()
            .getByResourceGroupWithResponse("rgopenapi", "exadbvmcluster1", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("key4195", "fakeTokenPlaceholder")).withZones(Arrays.asList("yd"))
            .withProperties(new ExadbVmClusterUpdateProperties().withNodeCount(17)).apply();
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
