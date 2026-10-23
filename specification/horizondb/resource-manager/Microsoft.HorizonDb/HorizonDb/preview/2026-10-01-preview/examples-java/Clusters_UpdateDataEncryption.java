
import com.azure.resourcemanager.horizondb.models.Cluster;
import com.azure.resourcemanager.horizondb.models.ClusterPropertiesForPatchUpdate;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterDataEncryptionForPatchUpdate;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_UpdateDataEncryption.json
     */
    /**
     * Sample code: Rotate the customer-managed key used by a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void rotateTheCustomerManagedKeyUsedByAHorizonDBCluster(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        Cluster resource = manager.clusters().getByResourceGroupWithResponse("example-resource-group",
            "example-cluster", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ClusterPropertiesForPatchUpdate().withDataEncryption(
            new HorizonDbClusterDataEncryptionForPatchUpdate().withPrimaryKeyUri("fakeTokenPlaceholder"))).apply();
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
