
import com.azure.resourcemanager.horizondb.models.BackupProperties;
import com.azure.resourcemanager.horizondb.models.ClusterProperties;
import com.azure.resourcemanager.horizondb.models.GeoRedundancyState;
import com.azure.resourcemanager.horizondb.models.PostgreSqlVersion;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_CreateOrReplaceMinimumDefaults.json
     */
    /**
     * Sample code: Create a PostgreSQL 18 HorizonDB cluster with minimum configuration and service defaults.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createAPostgreSQL18HorizonDBClusterWithMinimumConfigurationAndServiceDefaults(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().define("minimum-defaults").withRegion("eastus2")
            .withExistingResourceGroup("example-resource-group")
            .withProperties(new ClusterProperties().withPostgreSqlVersion(PostgreSqlVersion.V18)
                .withBackup(new BackupProperties().withGeoRedundancy(GeoRedundancyState.DISABLED)))
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
