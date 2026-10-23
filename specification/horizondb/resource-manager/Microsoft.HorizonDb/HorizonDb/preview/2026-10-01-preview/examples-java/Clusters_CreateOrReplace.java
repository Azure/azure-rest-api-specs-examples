
import com.azure.resourcemanager.horizondb.models.BackupProperties;
import com.azure.resourcemanager.horizondb.models.ClusterProperties;
import com.azure.resourcemanager.horizondb.models.GeoRedundancyState;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterIdentityType;
import com.azure.resourcemanager.horizondb.models.HorizonDbClusterManagedServiceIdentity;
import com.azure.resourcemanager.horizondb.models.Network;
import com.azure.resourcemanager.horizondb.models.PostgreSqlVersion;
import com.azure.resourcemanager.horizondb.models.PublicNetworkAccessState;
import com.azure.resourcemanager.horizondb.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Clusters CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Clusters_CreateOrReplace.json
     */
    /**
     * Sample code: Create or replace a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void createOrReplaceAHorizonDBCluster(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.clusters().define("example-cluster").withRegion("eastus2")
            .withExistingResourceGroup("example-resource-group").withTags(mapOf("environment", "production"))
            .withProperties(new ClusterProperties().withPostgreSqlVersion(PostgreSqlVersion.V17)
                .withNetwork(new Network().withPublicAccess(PublicNetworkAccessState.ENABLED))
                .withBackup(new BackupProperties().withGeoRedundancy(GeoRedundancyState.ENABLED).withRetentionDays(14)))
            .withIdentity(new HorizonDbClusterManagedServiceIdentity()
                .withType(HorizonDbClusterIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/identity-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/fabric-writer",
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
