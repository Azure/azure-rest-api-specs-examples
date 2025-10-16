
import com.azure.resourcemanager.oracledatabase.models.CloudVmCluster;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for CloudVmClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_patch.json
     */
    /**
     * Sample code: CloudVmClusters_update.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudVmClustersUpdate(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        CloudVmCluster resource = manager.cloudVmClusters()
            .getByResourceGroupWithResponse("rg000", "cluster1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
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
